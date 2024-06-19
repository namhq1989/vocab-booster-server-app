package infrastructure

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/caching"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type CachingRepository struct {
	caching caching.Operations
	ttl     time.Duration
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching: caching,
		ttl:     1 * time.Hour, // 1 hour
	}
}

func (r CachingRepository) GetUserSubscriptionPlan(ctx *appcontext.AppContext, userID string) (*domain.Plan, error) {
	key := r.generateUserSubscriptionKey(userID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	plan, ok := domain.SubscriptionPlans[dataStr]
	if !ok {
		return nil, apperrors.Subscription.InvalidPlan
	}

	dPlan := domain.ToPlan(plan.ID)
	if !dPlan.IsValid() {
		return nil, apperrors.Subscription.InvalidPlan
	}

	return &dPlan, nil
}

func (r CachingRepository) SetUserSubscriptionPlan(ctx *appcontext.AppContext, userID string, plan string) error {
	key := r.generateUserSubscriptionKey(userID)
	r.caching.SetTTL(ctx, key, plan, r.ttl)
	return nil
}

func (r CachingRepository) generateUserSubscriptionKey(userID string) string {
	return r.caching.GenerateKey("subscription", userID)
}
