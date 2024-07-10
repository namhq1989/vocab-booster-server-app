package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetSubscriptionPlansHandler struct{}

func NewGetSubscriptionPlansHandler() GetSubscriptionPlansHandler {
	return GetSubscriptionPlansHandler{}
}

func (GetSubscriptionPlansHandler) GetSubscriptionPlans(ctx *appcontext.AppContext, performerID string, _ dto.GetSubscriptionPlansRequest) (*dto.GetSubscriptionPlansResponse, error) {
	ctx.Logger().Info("[query] new get subscription plans request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("done get subscription plans request")
	return &dto.GetSubscriptionPlansResponse{
		Plans: []dto.SubscriptionPlan{
			dto.SubscriptionPlan{}.FromDomain(domain.SubscriptionPlans[domain.PlanFree.String()]),
			dto.SubscriptionPlan{}.FromDomain(domain.SubscriptionPlans[domain.PlanPremiumMonthly.String()]),
			dto.SubscriptionPlan{}.FromDomain(domain.SubscriptionPlans[domain.PlanPremiumQuarterly.String()]),
			dto.SubscriptionPlan{}.FromDomain(domain.SubscriptionPlans[domain.PlanPremiumYearly.String()]),
		},
	}, nil
}
