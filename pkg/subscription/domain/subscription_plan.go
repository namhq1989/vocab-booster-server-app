package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
)

type SubscriptionPlan struct {
	ID              string
	Price           float64
	DiscountPercent float64
	FinalPrice      float64
	Duration        time.Duration
}

var SubscriptionPlans = map[string]SubscriptionPlan{
	PlanFree.String(): {
		ID:              PlanFree.String(),
		Price:           0,
		DiscountPercent: 0,
		FinalPrice:      0,
		Duration:        0,
	},
	PlanPremiumMonthly.String(): {
		ID:              PlanPremiumMonthly.String(),
		Price:           40000,
		DiscountPercent: 0,
		FinalPrice:      40000,
		Duration:        30 * 24 * time.Hour,
	},
	PlanPremiumQuarterly.String(): {
		ID:              PlanPremiumQuarterly.String(),
		Price:           120000,
		DiscountPercent: 20,
		FinalPrice:      96000,
		Duration:        90 * 24 * time.Hour,
	},
	PlanPremiumYearly.String(): {
		ID:              PlanPremiumYearly.String(),
		Price:           480000,
		DiscountPercent: 50,
		FinalPrice:      240000,
		Duration:        365 * 24 * time.Hour,
	},
}

func GetSubscriptionPlan(planID string) (*SubscriptionPlan, error) {
	plan, ok := SubscriptionPlans[planID]
	if !ok {
		return nil, apperrors.Subscription.InvalidPlan
	}
	return &plan, nil
}

func (d SubscriptionPlan) IsFree() bool {
	return d.ID == PlanFree.String()
}
