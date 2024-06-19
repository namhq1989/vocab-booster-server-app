package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type CreateUserSubscriptionHandler struct {
	userSubscriptionHub domain.UserSubscriptionHub
}

func NewCreateUserSubscriptionHandler(userSubscriptionHub domain.UserSubscriptionHub) CreateUserSubscriptionHandler {
	return CreateUserSubscriptionHandler{
		userSubscriptionHub: userSubscriptionHub,
	}
}

func (h CreateUserSubscriptionHandler) CreateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error) {
	ctx.Logger().Info("[hub] new create user subscription request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("find subscription by user id in db")
	us, err := h.userSubscriptionHub.FindUserSubscriptionByUserID(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find subscription by user id in db", err, appcontext.Fields{})
		return nil, err
	}
	if us != nil {
		ctx.Logger().ErrorText("user subscription already exists")
		return nil, apperrors.Common.AlreadyExisted
	}

	ctx.Logger().Text("create new subscription's model")
	us, err = domain.NewUserSubscription(req.GetUserId(), domain.PlanFree.String())
	if err != nil {
		ctx.Logger().Error("failed to create new subscription's model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist subscription in db")
	err = h.userSubscriptionHub.CreateUserSubscription(ctx, *us)
	if err != nil {
		ctx.Logger().Error("failed to persist subscription in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create user subscription request")
	return &subscriptionpb.CreateUserSubscriptionResponse{
		Id: us.ID,
	}, nil
}
