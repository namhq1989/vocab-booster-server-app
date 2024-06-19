package apperrors

import "errors"

var Subscription = struct {
	SubscriptionNotFound  error
	InvalidSubscriptionID error
	InvalidPlan           error
}{
	SubscriptionNotFound:  errors.New("subscription_not_found"),
	InvalidSubscriptionID: errors.New("subscription_invalid_id"),
	InvalidPlan:           errors.New("subscription_invalid_plan"),
}
