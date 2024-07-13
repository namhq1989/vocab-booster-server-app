package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type QueueRepository struct {
	queue queue.Operations
}

func NewQueueRepository(queue queue.Operations) QueueRepository {
	return QueueRepository{
		queue: queue,
	}
}

func (r QueueRepository) ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.ExerciseAnswered, payload, -1)
}

func (r QueueRepository) AddAnswerExercisePoint(ctx *appcontext.AppContext, payload domain.QueueAddAnswerExercisePoint) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.GamificationAddAnswerExercisePoint, payload, -1)
}
