package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type CanPerformActionHandler struct {
	service domain.Service
}

func NewCanPerformActionHandler(
	service domain.Service,
) CanPerformActionHandler {
	return CanPerformActionHandler{
		service: service,
	}
}

func (h CanPerformActionHandler) CanPerformAction(ctx *appcontext.AppContext, req *subscriptionpb.CanPerformActionRequest) (*subscriptionpb.CanPerformActionResponse, error) {
	ctx.Logger().Info("[hub] new can perform action request", appcontext.Fields{"userID": req.GetUserId(), "action": req.GetAction(), "totalPerformedToday": req.GetTotalPerformedToday()})

	ctx.Logger().Text("check action value")
	action := domain.ToAction(req.GetAction())
	if !action.IsValid() {
		ctx.Logger().ErrorText("invalid action")
		return nil, apperrors.Common.InvalidAction
	}

	ctx.Logger().Text("get user subscription")
	us, err := h.service.GetUserSubscription(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to get user subscription", err, appcontext.Fields{})
		return &subscriptionpb.CanPerformActionResponse{Can: false}, nil
	}

	ctx.Logger().Info("check action limitation", appcontext.Fields{"plan": us.Plan.String()})
	hasExceededLimit := action.HasExceededLimit(us.Plan.String(), int(req.GetTotalPerformedToday()))

	ctx.Logger().Text("done can perform action request")
	return &subscriptionpb.CanPerformActionResponse{Can: !hasExceededLimit}, nil
}
