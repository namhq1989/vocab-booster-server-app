package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type FindUserSubscriptionHandler struct {
	userSubscriptionHub domain.UserSubscriptionHub
}

func NewFindUserSubscriptionHandler(userSubscriptionHub domain.UserSubscriptionHub) FindUserSubscriptionHandler {
	return FindUserSubscriptionHandler{
		userSubscriptionHub: userSubscriptionHub,
	}
}

func (h FindUserSubscriptionHandler) FindUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.FindUserSubscriptionRequest) (*subscriptionpb.FindUserSubscriptionResponse, error) {
	ctx.Logger().Info("[hub] new get user subscription request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("find subscription by user id in db")
	us, err := h.userSubscriptionHub.FindUserSubscriptionByUserID(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find subscription by user id in db", err, appcontext.Fields{})
		return nil, err
	}
	if us == nil {
		ctx.Logger().ErrorText("user subscription not found")
		return nil, apperrors.Subscription.SubscriptionNotFound
	}

	ctx.Logger().Text("done get user subscription request")
	return &subscriptionpb.FindUserSubscriptionResponse{
		Plan: &subscriptionpb.SubscriptionPlan{
			Id:      us.ID,
			Plan:    us.Plan.String(),
			StartAt: manipulation.ConvertToProtoTimestamp(us.StartAt),
			EndAt:   manipulation.ConvertToProtoTimestamp(us.EndAt),
		},
	}, nil
}
