package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/caching"
	"github.com/namhq1989/vocab-booster-server-app/internal/config"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/monitoring"
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/waiter"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Caching() *caching.Caching
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	JWT() *appjwt.JWT
	Monitoring() *monitoring.Monitoring
	Queue() *queue.Queue
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
