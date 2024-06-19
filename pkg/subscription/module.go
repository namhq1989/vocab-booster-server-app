package subscription

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/rest"
)

type Module struct{}

func (Module) Name() string {
	return "SUBSCRIPTION"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		_ = infrastructure.NewUserSubscriptionRepository(mono.Database())

		userSubscriptionHub = infrastructure.NewUserSubscriptionHub(mono.Database())

		// app
		app = application.New()
		hub = grpc.New(userSubscriptionHub)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), hub); err != nil {
		return err
	}

	return nil
}