package domain

import "github.com/namhq1989/vocab-booster-server-app/core/appcontext"

type UserSubscriptionHub interface {
	FindUserSubscriptionByUserID(ctx *appcontext.AppContext, userID string) (*UserSubscription, error)
	CreateUserSubscription(ctx *appcontext.AppContext, us UserSubscription) error
}
