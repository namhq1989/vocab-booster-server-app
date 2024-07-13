package dbmodel

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPoint struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"userId"`
	Point  int64              `bson:"point"`
}

func (m UserPoint) ToDomain() domain.UserPoint {
	return domain.UserPoint{
		ID:     m.ID.Hex(),
		UserID: m.UserID.Hex(),
		Point:  m.Point,
	}
}

func (UserPoint) FromDomain(point domain.UserPoint) (*UserPoint, error) {
	id, err := database.ObjectIDFromString(point.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(point.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserPoint{
		ID:     id,
		UserID: uid,
		Point:  point.Point,
	}, nil
}
