package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Handlers interface {
		ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		ExerciseAnsweredHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	queueRepository domain.QueueRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			ExerciseAnsweredHandler: NewExerciseAnsweredHandler(queueRepository),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.ExerciseAnswered), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueExerciseAnsweredPayload](bgCtx, t, queue.ParsePayload[domain.QueueExerciseAnsweredPayload], w.ExerciseAnswered)
	})
}
