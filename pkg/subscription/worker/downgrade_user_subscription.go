package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type DowngradeUserSubscriptionHandler struct {
	userSubscriptionRepository domain.UserSubscriptionRepository
}

func NewDowngradeUserSubscriptionHandler(
	userSubscriptionRepository domain.UserSubscriptionRepository,
) DowngradeUserSubscriptionHandler {
	return DowngradeUserSubscriptionHandler{
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (w DowngradeUserSubscriptionHandler) DowngradeUserSubscription(ctx *appcontext.AppContext, payload domain.QueueDowngradeUserSubscriptionPayload) error {
	ctx.Logger().Info("start downgrading user subscription", appcontext.Fields{"userID": payload.Subscription.UserID})

	ctx.Logger().Text("set data")
	us := payload.Subscription
	if err := us.DowngradeToFreePlan(); err != nil {
		ctx.Logger().Error("failed to set data", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update user subscription in db")
	if err := w.userSubscriptionRepository.UpsertUserSubscription(ctx, us); err != nil {
		ctx.Logger().Error("failed to update user subscription in db", err, appcontext.Fields{})
		return err
	}

	return nil
}
