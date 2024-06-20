package domain

import "github.com/namhq1989/vocab-booster-server-app/core/appcontext"

type CachingRepository interface {
	GetUserSubscription(ctx *appcontext.AppContext, userID string) (*UserSubscription, error)
	SetUserSubscription(ctx *appcontext.AppContext, userID string, plan UserSubscription) error
}
