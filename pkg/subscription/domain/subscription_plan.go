package domain

import "time"

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
