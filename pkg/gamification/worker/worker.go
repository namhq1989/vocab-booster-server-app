package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Handlers interface {
		AddAnswerExercisePoint(ctx *appcontext.AppContext, payload domain.QueueAddAnswerExercisePoint) error
		AddContributeVocabularySentencePoint(ctx *appcontext.AppContext, payload domain.QueueAddContributeVocabularySentencePoint) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		AddAnswerExercisePointHandler
		AddContributeVocabularySentencePointHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	pointRepository domain.PointRepository,
	userPointRepository domain.UserPointRepository,
	service domain.Service,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			AddAnswerExercisePointHandler:               NewAddAnswerExercisePointHandler(pointRepository, userPointRepository, service),
			AddContributeVocabularySentencePointHandler: NewAddContributeVocabularySentencePointHandler(pointRepository, userPointRepository, service),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GamificationAddAnswerExercisePoint), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueAddAnswerExercisePoint](bgCtx, t, queue.ParsePayload[domain.QueueAddAnswerExercisePoint], w.AddAnswerExercisePoint)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GamificationAddContributeVocabularySentencePoint), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueAddContributeVocabularySentencePoint](bgCtx, t, queue.ParsePayload[domain.QueueAddContributeVocabularySentencePoint], w.AddContributeVocabularySentencePoint)
	})
}
