package domain

import "github.com/namhq1989/vocab-booster-server-app/core/appcontext"

type Service interface {
	GetUserSubscription(ctx *appcontext.AppContext, userID string) (*UserSubscription, error)
}
