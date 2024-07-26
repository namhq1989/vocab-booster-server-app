package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-utilities/language"
)

type Exercise struct {
	ID            string
	Audio         string
	Level         string
	Content       language.Multilingual
	Vocabulary    string
	CorrectAnswer string
	Options       []string
	CorrectStreak int
	IsFavorite    bool
	IsMastered    bool
	UpdatedAt     time.Time
	NextReviewAt  time.Time
}

//
// AGGREGATED EXERCISES
//

type UserAggregatedExercise struct {
	Date     string
	Exercise int64
}
