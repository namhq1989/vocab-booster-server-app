package shared

import (
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) GetUserSubscription(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	ctx.Logger().Text("get user subscription plan in caching layer first")
	us, err := s.cachingRepository.GetUserSubscription(ctx, userID)
	if us != nil {
		ctx.Logger().Text("found plan in caching layer")
		return us, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to get user subscription plan in caching layer", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find user subscription in db")
	us, err = s.userSubscriptionRepository.FindUserSubscriptionByUserID(ctx, userID)
	if err != nil {
		ctx.Logger().Error("failed to find user subscription in db", err, appcontext.Fields{})
		return nil, apperrors.Common.InvalidAction
	}
	if us == nil {
		ctx.Logger().ErrorText("user subscription not found, create new one")
		us, err = domain.NewUserSubscription(userID, domain.PlanFree.String())
		if err != nil {
			ctx.Logger().Error("failed to create new user subscription", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist user subscription in db")
		err = s.userSubscriptionRepository.UpsertUserSubscription(ctx, *us)
		if err != nil {
			ctx.Logger().Error("failed to create user subscription", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("done get user subscription plan")
	return us, nil
}
