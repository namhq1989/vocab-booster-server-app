package main

import (
	"crypto/subtle"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/vocab-booster-server-app/internal/caching"
	"github.com/namhq1989/vocab-booster-server-app/internal/config"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/monitoring"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/waiter"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user"
	"github.com/namhq1989/vocab-booster-utilities/logger"
)

func main() {
	var err error

	// config
	cfg := config.Init()

	// logger
	logger.Init(cfg.Environment)

	// app error
	apperrors.Init()

	// server
	a := app{}
	a.cfg = cfg

	// rest
	a.rest = initRest(cfg)

	// grpc
	a.rpc = initRPC()

	// jwt
	a.jwt, err = appjwt.Init(cfg.AccessTokenSecret, time.Second*time.Duration(cfg.AccessTokenTTL))
	if err != nil {
		panic(err)
	}

	// database
	a.database = database.NewDatabaseClient(cfg.MongoURL, cfg.MongoDBName)

	// queue
	a.queue = queue.Init(cfg.QueueRedisURL, cfg.QueueConcurrency)

	// init queue's dashboard
	a.rest.Any(fmt.Sprintf("%s/*", queue.DashboardPath), echo.WrapHandler(queue.EnableDashboard(cfg.QueueRedisURL)), middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if !cfg.IsEnvRelease {
			return true, nil
		}
		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// caching
	a.caching = caching.NewCachingClient(cfg.CachingRedisURL)

	// monitoring
	a.monitoring = monitoring.Init(a.Rest(), cfg.SentryDSN, cfg.SentryMachine, cfg.Environment)

	// waiter
	a.waiter = waiter.New(waiter.CatchSignals())

	// modules
	a.modules = []monolith.Module{
		&auth.Module{},
		&user.Module{},
		&subscription.Module{},
	}

	// start
	if err = a.startupModules(); err != nil {
		panic(err)
	}

	fmt.Println("--- started vocab-booster-server-app application")
	defer fmt.Println("--- stopped vocab-booster-server-app application")

	// wait for other service starts
	a.waiter.Add(
		a.waitForRest,
		a.waitForRPC,
	)
	if err = a.waiter.Wait(); err != nil {
		panic(err)
	}
}
