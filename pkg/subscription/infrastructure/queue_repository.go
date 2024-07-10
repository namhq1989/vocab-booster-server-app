package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
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

func (r QueueRepository) DowngradeUserSubscription(ctx *appcontext.AppContext, payload domain.QueueDowngradeUserSubscriptionPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.DowngradeUserSubscription, payload, -1)
}
