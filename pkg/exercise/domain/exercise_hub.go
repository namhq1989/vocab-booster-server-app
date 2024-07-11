package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseHub interface {
	GetExercises(ctx *appcontext.AppContext, userID, lang string) ([]Exercise, error)
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
