package user

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/grpcclient"
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
	gamificationGRPCClient, err := grpcclient.NewGamificationClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	exerciseGRPCClient, err := grpcclient.NewExerciseClient(ctx, mono.Config().EndpointExerciseGrpc)
	if err != nil {
		return err
	}

	var (
		userRepository = infrastructure.NewUserRepository(mono.Database())

		userHub         = infrastructure.NewUserHub(mono.Database())
		gamificationHub = infrastructure.NewGamificationHub(gamificationGRPCClient)
		exerciseHub     = infrastructure.NewExerciseHub(exerciseGRPCClient)

		// app
		app = application.New(userRepository, gamificationHub, exerciseHub)
		hub = grpc.New(userHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// grpc server
	if err = grpc.RegisterServer(ctx, mono.RPC(), hub); err != nil {
		return err
	}

	return nil
}
