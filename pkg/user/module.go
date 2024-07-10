package user

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/grpc"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/rest"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "USER"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		userRepository    = infrastructure.NewUserRepository(mono.Database())
		journeyRepository = infrastructure.NewJourneyRepository(mono.Database())

		userHub = infrastructure.NewUserHub(mono.Database())

		// app
		app = application.New(userRepository, journeyRepository)
		hub = grpc.New(userHub)
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
