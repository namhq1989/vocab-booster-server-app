package domain

import "github.com/namhq1989/vocab-booster-server-app/core/appcontext"

type CachingRepository interface {
	GetUserSubscriptionPlan(ctx *appcontext.AppContext, userID string) (*Plan, error)
	SetUserSubscriptionPlan(ctx *appcontext.AppContext, userID, plan string) error
}
