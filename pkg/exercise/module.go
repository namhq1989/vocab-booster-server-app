package exercise

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/rest"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/worker"
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

	gamificationGRPCClient, err := grpcclient.NewGamificationClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		queueRepository = infrastructure.NewQueueRepository(mono.Queue())

		exerciseHub     = infrastructure.NewExerciseHub(exerciseGRPCClient)
		gamificationHub = infrastructure.NewGamificationHub(gamificationGRPCClient)

		// app
		app = application.New(queueRepository, exerciseHub, gamificationHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		queueRepository,
	)
	w.Start()

	return nil
}
