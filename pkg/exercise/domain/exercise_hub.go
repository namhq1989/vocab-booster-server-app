package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseHub interface {
	AnswerExercise(ctx *appcontext.AppContext, payload AnswerExercisePayload) (*AnswerExerciseResult, error)
	GetExercises(ctx *appcontext.AppContext, userID, lang, collectionID string) ([]Exercise, error)
	GetReadyForReviewExercises(ctx *appcontext.AppContext, userID, lang, timezone string) ([]Exercise, error)
	GetExerciseCollections(ctx *appcontext.AppContext, userID, lang string) ([]ExerciseCollection, error)
	AggregateUserExercisesInTimeRange(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]UserAggregatedExercise, error)
	ChangeExerciseFavorite(ctx *appcontext.AppContext, userID, exerciseID string, isFavorite bool) (bool, error)
}

type AnswerExercisePayload struct {
	UserID     string
	ExerciseID string
	IsCorrect  bool
	Timezone   string
}

func NewAnswerExercisePayload(userID, exerciseID string, isCorrect bool, timezone string) (*AnswerExercisePayload, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidObjectID(exerciseID) {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	return &AnswerExercisePayload{
		UserID:     userID,
		ExerciseID: exerciseID,
		IsCorrect:  isCorrect,
		Timezone:   timezone,
	}, nil
}

type AnswerExerciseResult struct {
	NextReviewAt time.Time
}
