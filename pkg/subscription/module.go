package subscription

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/rest"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/shared"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/worker"
)

type Module struct{}

func (Module) Name() string {
	return "SUBSCRIPTION"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		userSubscriptionRepository        = infrastructure.NewUserSubscriptionRepository(mono.Database())
		userSubscriptionHistoryRepository = infrastructure.NewUserSubscriptionHistoryRepository(mono.Database())

		cachingRepository = infrastructure.NewCachingRepository(mono.Caching())
		queueRepository   = infrastructure.NewQueueRepository(mono.Queue())

		userSubscriptionHub = infrastructure.NewUserSubscriptionHub(mono.Database())

		service = shared.NewService(userSubscriptionRepository, cachingRepository)

		// app
		app = application.New()
		hub = grpc.New(userSubscriptionRepository, userSubscriptionHistoryRepository, userSubscriptionHub, service)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), hub); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		userSubscriptionRepository,
		queueRepository,
	)
	w.Start()

	return nil
}
