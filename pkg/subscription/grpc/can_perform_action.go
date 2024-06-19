package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type CanPerformActionHandler struct {
	userSubscriptionRepository domain.UserSubscriptionRepository
	cachingRepository          domain.CachingRepository
}

func NewCanPerformActionHandler(
	userSubscriptionRepository domain.UserSubscriptionRepository,
	cachingRepository domain.CachingRepository,
) CanPerformActionHandler {
	return CanPerformActionHandler{
		userSubscriptionRepository: userSubscriptionRepository,
		cachingRepository:          cachingRepository,
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

	ctx.Logger().Text("get user subscription plan")
	plan, _ := h.getUserSubscriptionPlan(ctx, req.GetUserId())

	ctx.Logger().Info("check action limitation", appcontext.Fields{"plan": plan.String()})
	hasExceededLimit := action.HasExceededLimit(plan.String(), int(req.GetTotalPerformedToday()))

	ctx.Logger().Text("done can perform action request")
	return &subscriptionpb.CanPerformActionResponse{Can: !hasExceededLimit}, nil
}

func (h CanPerformActionHandler) getUserSubscriptionPlan(ctx *appcontext.AppContext, userID string) (domain.Plan, error) {
	ctx.Logger().Text("get user subscription plan in caching layer first")
	plan, err := h.cachingRepository.GetUserSubscriptionPlan(ctx, userID)
	if plan != nil {
		ctx.Logger().Text("found plan in caching layer")
		return *plan, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to get user subscription plan in caching layer", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find user subscription in db")
	us, err := h.userSubscriptionRepository.FindUserSubscriptionByUserID(ctx, userID)
	if err != nil {
		ctx.Logger().Error("failed to find user subscription in db", err, appcontext.Fields{})
		return domain.PlanFree, nil
	}

	ctx.Logger().Text("done get user subscription plan")
	return us.Plan, nil
}
