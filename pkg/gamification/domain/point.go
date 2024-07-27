package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type PointRepository interface {
	CreatePoint(ctx *appcontext.AppContext, point Point) error
	AggregateUserPointsInTimeRange(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]UserAggregatedPoint, error)
}

type Point struct {
	ID           string
	UserID       string
	Action       Action
	ExerciseID   string
	VocabularyID string
	Point        int64
	CreatedAt    time.Time
}

func NewPoint(userID, exerciseID, vocabularyID, action string, point int64) (*Point, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if exerciseID == "" && vocabularyID == "" {
		return nil, apperrors.Gamification.InvalidPointData
	}

	if exerciseID != "" && !database.IsValidObjectID(exerciseID) {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	if vocabularyID != "" && !database.IsValidObjectID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	if point < 0 || point > 100 {
		return nil, apperrors.Gamification.InvalidPoint
	}

	dAction := ToAction(action)
	if !dAction.IsValid() {
		return nil, apperrors.Common.InvalidAction
	}

	if dAction.IsAnswerExercise() && exerciseID == "" {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	if dAction.IsContributeVocabularySentence() && vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &Point{
		ID:           database.NewStringID(),
		UserID:       userID,
		Action:       dAction,
		ExerciseID:   exerciseID,
		VocabularyID: vocabularyID,
		Point:        point,
		CreatedAt:    manipulation.NowUTC(),
	}, nil
}

//
// AGGREGATED POINTS
//

type UserAggregatedPoint struct {
	Date  string
	Point int64
}
