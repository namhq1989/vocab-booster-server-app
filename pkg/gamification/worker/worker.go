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
		ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPoint) error
		VocabularySentenceContributed(ctx *appcontext.AppContext, payload domain.QueueVocabularySentenceContributedPoint) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		ExerciseAnsweredHandler
		VocabularySentenceContributedHandler
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
	completionTimeRepository domain.CompletionTimeRepository,
	userStatsRepository domain.UserStatsRepository,
	service domain.Service,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			ExerciseAnsweredHandler:              NewExerciseAnsweredHandler(pointRepository, completionTimeRepository, userStatsRepository, service),
			VocabularySentenceContributedHandler: NewVocabularySentenceContributedHandler(pointRepository, completionTimeRepository, userStatsRepository, service),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GamificationExerciseAnswered), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueExerciseAnsweredPoint](bgCtx, t, queue.ParsePayload[domain.QueueExerciseAnsweredPoint], w.ExerciseAnswered)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GamificationVocabularySentenceContributed), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueVocabularySentenceContributedPoint](bgCtx, t, queue.ParsePayload[domain.QueueVocabularySentenceContributedPoint], w.VocabularySentenceContributed)
	})
}
