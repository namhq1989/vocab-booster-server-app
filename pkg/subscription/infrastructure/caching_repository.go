package infrastructure

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"

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

func (r CachingRepository) GetUserSubscription(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	key := r.generateUserSubscriptionKey(userID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var us *domain.UserSubscription
	if err = json.Unmarshal([]byte(dataStr), &us); err != nil {
		_, _ = r.caching.Del(ctx, key)
		return nil, nil
	}

	return us, nil
}

func (r CachingRepository) SetUserSubscription(ctx *appcontext.AppContext, userID string, us domain.UserSubscription) error {
	key := r.generateUserSubscriptionKey(userID)
	r.caching.SetTTL(ctx, key, us, r.ttl)
	return nil
}

func (r CachingRepository) generateUserSubscriptionKey(userID string) string {
	return r.caching.GenerateKey("subscription", fmt.Sprintf("user_subscription_plan_%s", userID))
}
