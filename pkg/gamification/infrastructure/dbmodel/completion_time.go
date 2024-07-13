package dbmodel

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompletionTime struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	Action    string             `bson:"action"`
	Seconds   int                `bson:"seconds"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (m CompletionTime) ToDomain() domain.CompletionTime {
	return domain.CompletionTime{
		ID:        m.ID.Hex(),
		UserID:    m.UserID.Hex(),
		Action:    domain.ToAction(m.Action),
		Seconds:   m.Seconds,
		CreatedAt: m.CreatedAt,
	}
}

func (CompletionTime) FromDomain(completionTime domain.CompletionTime) (*CompletionTime, error) {
	id, err := database.ObjectIDFromString(completionTime.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	userID, err := database.ObjectIDFromString(completionTime.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &CompletionTime{
		ID:        id,
		UserID:    userID,
		Action:    completionTime.Action.String(),
		Seconds:   completionTime.Seconds,
		CreatedAt: completionTime.CreatedAt,
	}, nil
}
