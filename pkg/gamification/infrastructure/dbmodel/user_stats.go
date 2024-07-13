package dbmodel

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStats struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserID         primitive.ObjectID `bson:"userId"`
	Point          int64              `bson:"point"`
	CompletionTime int                `bson:"completionTime"`
}

func (m UserStats) ToDomain() domain.UserStats {
	return domain.UserStats{
		ID:             m.ID.Hex(),
		UserID:         m.UserID.Hex(),
		Point:          m.Point,
		CompletionTime: m.CompletionTime,
	}
}

func (UserStats) FromDomain(stats domain.UserStats) (*UserStats, error) {
	id, err := database.ObjectIDFromString(stats.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(stats.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserStats{
		ID:             id,
		UserID:         uid,
		Point:          stats.Point,
		CompletionTime: stats.CompletionTime,
	}, nil
}
