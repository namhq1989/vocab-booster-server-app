package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	GetUserSubscription(ctx *appcontext.AppContext, userID string) (*UserSubscription, error)
}
