package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type UpdateUserSubscriptionHandler struct {
	userSubscriptionRepository        domain.UserSubscriptionRepository
	userSubscriptionHistoryRepository domain.UserSubscriptionHistoryRepository
}

func NewUpdateUserSubscriptionHandler(
	userSubscriptionRepository domain.UserSubscriptionRepository,
	userSubscriptionHistoryRepository domain.UserSubscriptionHistoryRepository,
) UpdateUserSubscriptionHandler {
	return UpdateUserSubscriptionHandler{
		userSubscriptionRepository:        userSubscriptionRepository,
		userSubscriptionHistoryRepository: userSubscriptionHistoryRepository,
	}
}

func (h UpdateUserSubscriptionHandler) UpdateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.UpdateUserSubscriptionRequest) (*subscriptionpb.UpdateUserSubscriptionResponse, error) {
	ctx.Logger().Info("[hub] new update user subscription request", appcontext.Fields{"userID": req.GetUserId(), "plan": req.GetPlan(), "paymentID": req.GetPaymentId()})

	ctx.Logger().Text("get plan detail")
	plan, err := domain.GetSubscriptionPlan(req.GetPlan())
	if err != nil {
		ctx.Logger().Error("failed to get plan detail", err, appcontext.Fields{})
		return nil, err
	}
	if plan == nil {
		ctx.Logger().ErrorText("invalid plan")
		return nil, apperrors.Subscription.InvalidPlan
	}
	if plan.IsFree() {
		ctx.Logger().ErrorText("cannot update user subscription to plan Free, please use method downgrade for this operation")
		return nil, apperrors.Subscription.InvalidPlan
	}

	ctx.Logger().Text("find user subscription in db")
	us, err := h.userSubscriptionRepository.FindUserSubscriptionByUserID(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find user subscription in db", err, appcontext.Fields{})
		return nil, err
	}
	if us == nil {
		ctx.Logger().ErrorText("user subscription not found, create new one")
		us, err = domain.NewUserSubscription(req.GetUserId(), domain.PlanFree.String())
		if err != nil {
			ctx.Logger().Error("failed to create new user subscription", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist user subscription in db")
		err = h.userSubscriptionRepository.UpsertUserSubscription(ctx, *us)
		if err != nil {
			ctx.Logger().Error("failed to persist user subscription in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	if us.Plan.IsFree() {
		ctx.Logger().Text("user's current plan is Free, upgrade to Premium")
		if err = us.UpgradeToPremium(plan.ID); err != nil {
			ctx.Logger().Error("failed to upgrade user subscription to Premium", err, appcontext.Fields{})
			return nil, err
		}
	} else {
		ctx.Logger().Text("user's current plan is already Premium, just extend the duration")
		if err = us.ExtendDuration(plan.Duration); err != nil {
			ctx.Logger().Error("failed to extend user subscription duration", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("persist user subscription in db")
	if err = h.userSubscriptionRepository.UpsertUserSubscription(ctx, *us); err != nil {
		ctx.Logger().Error("failed to persist user subscription in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("create user subscription history")
	_ = h.createHistory(ctx, *us, req.GetPaymentId())

	ctx.Logger().Text("done update user subscription request")
	return &subscriptionpb.UpdateUserSubscriptionResponse{
		Id: us.ID,
	}, nil
}

func (h UpdateUserSubscriptionHandler) createHistory(ctx *appcontext.AppContext, us domain.UserSubscription, paymentID string) error {
	ctx.Logger().Text("create user subscription history")
	history, err := domain.NewUserSubscriptionHistory(us.UserID, paymentID)
	if err != nil {
		ctx.Logger().Error("failed to create user subscription history", err, appcontext.Fields{})
	} else {
		ctx.Logger().Text("persist user subscription history in db")
		err = h.userSubscriptionHistoryRepository.CreateUserSubscriptionHistory(ctx, *history)
		if err != nil {
			ctx.Logger().Error("failed to persist user subscription history in db", err, appcontext.Fields{})
		}
	}

	return nil
}
