package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type SubscriptionPlan struct {
	ID              string  `json:"id"`
	Price           float64 `json:"price"`
	DiscountPercent float64 `json:"discountPercent"`
	FinalPrice      float64 `json:"finalPrice"`
}

func (SubscriptionPlan) FromDomain(plan domain.SubscriptionPlan) SubscriptionPlan {
	return SubscriptionPlan{
		ID:              plan.ID,
		Price:           plan.Price,
		DiscountPercent: plan.DiscountPercent,
		FinalPrice:      plan.FinalPrice,
	}
}
