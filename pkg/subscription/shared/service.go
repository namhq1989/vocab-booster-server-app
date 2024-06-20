package shared

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type Service struct {
	userSubscriptionRepository domain.UserSubscriptionRepository
	cachingRepository          domain.CachingRepository
}

func NewService(
	userSubscriptionRepository domain.UserSubscriptionRepository,
	cachingRepository domain.CachingRepository,
) Service {
	return Service{
		userSubscriptionRepository: userSubscriptionRepository,
		cachingRepository:          cachingRepository,
	}
}
