package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserSubscriptionHandler struct {
	userSubscriptionHub domain.UserSubscriptionHub
}

func NewGetUserSubscriptionHandler(userSubscriptionHub domain.UserSubscriptionHub) GetUserSubscriptionHandler {
	return GetUserSubscriptionHandler{
		userSubscriptionHub: userSubscriptionHub,
	}
}

func (h GetUserSubscriptionHandler) GetUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.GetUserSubscriptionRequest) (*subscriptionpb.GetUserSubscriptionResponse, error) {
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
	return &subscriptionpb.GetUserSubscriptionResponse{
		Plan: &subscriptionpb.SubscriptionPlan{
			Id:      us.ID,
			Plan:    us.Plan.String(),
			StartAt: manipulation.ConvertToProtoTimestamp(us.StartAt),
			EndAt:   manipulation.ConvertToProtoTimestamp(us.EndAt),
		},
	}, nil
}
