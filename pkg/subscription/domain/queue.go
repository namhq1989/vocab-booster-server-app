package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type QueueRepository interface {
	DowngradeUserSubscription(ctx *appcontext.AppContext, payload QueueDowngradeUserSubscriptionPayload) error
}

type QueueScanExpiredUserSubscription struct{}

type QueueDowngradeUserSubscriptionPayload struct {
	Subscription UserSubscription
}
