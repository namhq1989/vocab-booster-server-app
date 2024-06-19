package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
)

type UserSubscriptionRepository interface {
	UpsertUserSubscription(ctx *appcontext.AppContext, subscription UserSubscription) error
}

type UserSubscription struct {
	ID      string
	UserID  string
	Plan    Plan
	StartAt time.Time
	EndAt   time.Time
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

	endAt := time.Now().AddDate(0, 0, int(subscriptionPlan.Duration.Hours()))
	endAt = manipulation.EndOfDate(endAt)

	return &UserSubscription{
		ID:      database.NewStringID(),
		UserID:  userID,
		Plan:    dPlan,
		StartAt: time.Now(),
		EndAt:   endAt,
	}, nil
}

func (d *UserSubscription) UpgradeToPremium(plan string) error {
	dPlan := ToPlan(plan)
	if !dPlan.IsValid() || !d.Plan.IsPremium() {
		return apperrors.Subscription.InvalidPlan
	}

	subscriptionPlan := SubscriptionPlans[plan]
	if err := d.ExtendDuration(subscriptionPlan.Duration); err != nil {
		return err
	}

	d.Plan = dPlan
	return nil
}

func (d *UserSubscription) ExtendDuration(duration time.Duration) error {
	if d.EndAt.IsZero() {
		d.EndAt = time.Now().AddDate(0, 0, int(duration.Hours()))
	} else {
		d.EndAt = d.EndAt.AddDate(0, 0, int(duration.Hours()))
	}
	d.EndAt = manipulation.EndOfDate(d.EndAt)
	return nil
}

func (d *UserSubscription) DowngradeToFreePlan() error {
	d.Plan = PlanFree
	d.StartAt = time.Time{}
	d.EndAt = time.Time{}
	return nil
}
