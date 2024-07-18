package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserSubscriptionRepository interface {
	FindUserSubscriptionByUserID(ctx *appcontext.AppContext, userID string) (*UserSubscription, error)
	FindExpiredUserSubscriptionsByDate(ctx *appcontext.AppContext, date time.Time) ([]UserSubscription, error)
	UpsertUserSubscription(ctx *appcontext.AppContext, subscription UserSubscription) error
}

type UserSubscription struct {
	ID        string
	UserID    string
	IsPremium bool
	Plan      Plan
	StartAt   time.Time
	EndAt     time.Time
}

func NewUserSubscription(userID string, plan string) (*UserSubscription, error) {
	if userID == "" {
		return nil, apperrors.User.InvalidUserID
	}

	dPlan := ToPlan(plan)
	if !dPlan.IsValid() {
		return nil, apperrors.Subscription.InvalidPlan
	}

	subscriptionPlan, ok := SubscriptionPlans[plan]
	if !ok {
		return nil, apperrors.Subscription.InvalidPlan
	}

	endAt := manipulation.Now().AddDate(0, 0, int(subscriptionPlan.Duration.Hours()))
	endAt = manipulation.EndOfDate(endAt)

	return &UserSubscription{
		ID:        database.NewStringID(),
		UserID:    userID,
		IsPremium: dPlan.IsPremium(),
		Plan:      dPlan,
		StartAt:   manipulation.Now(),
		EndAt:     endAt,
	}, nil
}

func (d *UserSubscription) UpgradeToPremium(plan string) error {
	dPlan := ToPlan(plan)
	if !dPlan.IsValid() || !dPlan.IsPremium() {
		return apperrors.Subscription.InvalidPlan
	}

	subscriptionPlan := SubscriptionPlans[plan]
	if err := d.ExtendDuration(subscriptionPlan.Duration); err != nil {
		return err
	}

	d.Plan = dPlan
	d.IsPremium = true
	return nil
}

func (d *UserSubscription) ExtendDuration(duration time.Duration) error {
	if d.EndAt.IsZero() {
		d.EndAt = manipulation.Now().AddDate(0, 0, int(duration.Hours()))
	} else {
		d.EndAt = d.EndAt.AddDate(0, 0, int(duration.Hours()))
	}
	d.EndAt = manipulation.EndOfDate(d.EndAt)
	return nil
}

func (d *UserSubscription) DowngradeToFreePlan() error {
	d.Plan = PlanFree
	d.IsPremium = false
	d.StartAt = time.Time{}
	d.EndAt = time.Time{}
	return nil
}
