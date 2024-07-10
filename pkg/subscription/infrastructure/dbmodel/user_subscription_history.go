package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSubscriptionHistory struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	PaymentID primitive.ObjectID `bson:"paymentId"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (m UserSubscriptionHistory) ToDomain() domain.UserSubscriptionHistory {
	return domain.UserSubscriptionHistory{
		ID:        m.ID.Hex(),
		UserID:    m.UserID.Hex(),
		PaymentID: m.PaymentID.Hex(),
		CreatedAt: m.CreatedAt,
	}
}

func (UserSubscriptionHistory) FromDomain(history domain.UserSubscriptionHistory) (*UserSubscriptionHistory, error) {
	id, err := database.ObjectIDFromString(history.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(history.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	pid, err := database.ObjectIDFromString(history.PaymentID)
	if err != nil {
		return nil, apperrors.Payment.InvalidPaymentID
	}

	return &UserSubscriptionHistory{
		ID:        id,
		UserID:    uid,
		PaymentID: pid,
		CreatedAt: history.CreatedAt,
	}, nil
}
