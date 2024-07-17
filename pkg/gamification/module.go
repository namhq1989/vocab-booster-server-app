package gamification

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/grpc"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/shared"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "GAMIFICATION"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		pointRepository          = infrastructure.NewPointRepository(mono.Database())
		completionTimeRepository = infrastructure.NewCompletionTimeRepository(mono.Database())
		userStatsRepository      = infrastructure.NewUserStatsRepository(mono.Database())
		service                  = shared.NewService(mono.Database(), pointRepository, completionTimeRepository, userStatsRepository)

		hub = grpc.New(userStatsRepository, pointRepository)
	)

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), hub); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		pointRepository,
		completionTimeRepository,
		userStatsRepository,
		service,
	)
	w.Start()

	return nil
}
