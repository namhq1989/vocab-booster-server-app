package exercise

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/rest"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "EXERCISE"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	exerciseGRPCClient, err := grpcclient.NewExerciseClient(ctx, mono.Config().EndpointExerciseGrpc)
	if err != nil {
		return err
	}

	var (
		exerciseHub = infrastructure.NewExerciseHub(exerciseGRPCClient)

		// app
		app = application.New(exerciseHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	return nil
}
