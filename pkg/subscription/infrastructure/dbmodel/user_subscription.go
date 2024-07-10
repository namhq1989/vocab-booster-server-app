package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSubscription struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	IsPremium bool               `bson:"isPremium"`
	Plan      string             `bson:"plan"`
	StartAt   time.Time          `bson:"startAt"`
	EndAt     time.Time          `bson:"endAt"`
}

func (m UserSubscription) ToDomain() domain.UserSubscription {
	return domain.UserSubscription{
		ID:        m.ID.Hex(),
		UserID:    m.UserID.Hex(),
		IsPremium: m.IsPremium,
		Plan:      domain.ToPlan(m.Plan),
		StartAt:   m.StartAt,
		EndAt:     m.EndAt,
	}
}

func (UserSubscription) FromDomain(us domain.UserSubscription) (*UserSubscription, error) {
	id, err := database.ObjectIDFromString(us.ID)
	if err != nil {
		return nil, apperrors.Subscription.InvalidSubscriptionID
	}

	uid, err := database.ObjectIDFromString(us.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserSubscription{
		ID:        id,
		UserID:    uid,
		IsPremium: us.IsPremium,
		Plan:      us.Plan.String(),
		StartAt:   us.StartAt,
		EndAt:     us.EndAt,
	}, nil
}
