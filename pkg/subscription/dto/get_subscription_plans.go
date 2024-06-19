package dto

type GetSubscriptionPlansRequest struct{}

type GetSubscriptionPlansResponse struct {
	Plans []SubscriptionPlan `json:"plans"`
}
