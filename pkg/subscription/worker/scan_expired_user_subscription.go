package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type ScanExpiredUserSubscriptionHandler struct {
	userSubscriptionRepository domain.UserSubscriptionRepository
	queueRepository            domain.QueueRepository
}

func NewScanExpiredUserSubscriptionHandler(
	userSubscriptionRepository domain.UserSubscriptionRepository,
	queueRepository domain.QueueRepository,
) ScanExpiredUserSubscriptionHandler {
	return ScanExpiredUserSubscriptionHandler{
		userSubscriptionRepository: userSubscriptionRepository,
		queueRepository:            queueRepository,
	}
}

func (w ScanExpiredUserSubscriptionHandler) ScanExpiredUserSubscription(ctx *appcontext.AppContext, _ domain.QueueScanExpiredUserSubscription) error {
	startOfToday := manipulation.StartOfToday()

	ctx.Logger().Info("[worker] find all docs that expired", appcontext.Fields{"date": startOfToday.String()})
	uss, err := w.userSubscriptionRepository.FindExpiredUserSubscriptionsByDate(ctx, startOfToday)
	if err != nil {
		ctx.Logger().Error("failed to find all docs that expired today", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Info("total docs that expired today", appcontext.Fields{"count": len(uss)})

	if len(uss) == 0 {
		return nil
	}

	for _, us := range uss {
		ctx.Logger().Text("add doc to queue for downgrade user subscription")
		if err = w.queueRepository.DowngradeUserSubscription(ctx, domain.QueueDowngradeUserSubscriptionPayload{
			Subscription: us,
		}); err != nil {
			ctx.Logger().Error("failed to create user subscription history", err, appcontext.Fields{})
		}
	}

	return nil
}
