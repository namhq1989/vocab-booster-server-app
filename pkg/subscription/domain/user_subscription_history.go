package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
)

type UserSubscriptionHistoryRepository interface {
	CreateUserSubscriptionHistory(ctx *appcontext.AppContext, history UserSubscriptionHistory) error
}

type UserSubscriptionHistory struct {
	ID        string
	UserID    string
	PaymentID string
	CreatedAt time.Time
}

func NewUserSubscriptionHistory(userID, paymentID string) (*UserSubscriptionHistory, error) {
	if userID == "" {
		return nil, apperrors.User.InvalidUserID
	}

	if paymentID == "" {
		return nil, apperrors.Payment.InvalidPaymentID
	}

	return &UserSubscriptionHistory{
		ID:        database.NewStringID(),
		UserID:    userID,
		PaymentID: paymentID,
		CreatedAt: time.Now(),
	}, nil
}
