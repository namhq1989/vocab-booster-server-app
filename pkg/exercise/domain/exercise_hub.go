package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseHub interface {
	AnswerExercise(ctx *appcontext.AppContext, payload AnswerExercisePayload) (*AnswerExerciseResult, error)
	GetExercises(ctx *appcontext.AppContext, userID, lang, level string) ([]Exercise, error)
	GetReadyForReviewExercises(ctx *appcontext.AppContext, userID, lang string) ([]Exercise, error)
}

type Exercise struct {
	ID            string
	Audio         string
	Level         string
	Content       string
	Translated    string
	Vocabulary    string
	CorrectAnswer string
	Options       []string
	CorrectStreak int
	IsFavorite    bool
	IsMastered    bool
	UpdatedAt     time.Time
	NextReviewAt  time.Time
}

type AnswerExercisePayload struct {
	UserID     string
	ExerciseID string
	IsCorrect  bool
}

func NewAnswerExercisePayload(userID, exerciseID string, isCorrect bool) (*AnswerExercisePayload, error) {
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
	}, nil
}

type AnswerExerciseResult struct {
	NextReviewAt time.Time
}
