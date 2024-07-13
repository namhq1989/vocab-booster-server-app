package dbmodel

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Point struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserID       primitive.ObjectID `bson:"userId"`
	Action       string             `bson:"action"`
	ExerciseID   string             `bson:"exerciseId"`
	VocabularyID string             `bson:"vocabulary"`
	Point        int64              `bson:"point"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

func (m Point) ToDomain() domain.Point {
	return domain.Point{
		ID:           m.ID.Hex(),
		UserID:       m.UserID.Hex(),
		Action:       domain.ToAction(m.Action),
		ExerciseID:   m.ExerciseID,
		VocabularyID: m.VocabularyID,
		Point:        m.Point,
		CreatedAt:    m.CreatedAt,
	}
}

func (Point) FromDomain(point domain.Point) (*Point, error) {
	id, err := database.ObjectIDFromString(point.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(point.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	if point.ExerciseID != "" && !database.IsValidObjectID(point.ExerciseID) {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	if point.VocabularyID != "" && !database.IsValidObjectID(point.VocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &Point{
		ID:           id,
		UserID:       uid,
		Action:       point.Action.String(),
		ExerciseID:   point.ExerciseID,
		VocabularyID: point.VocabularyID,
		Point:        point.Point,
		CreatedAt:    point.CreatedAt,
	}, nil
}
