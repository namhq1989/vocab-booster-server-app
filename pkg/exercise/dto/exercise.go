package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type Exercise struct {
	ID            string                    `json:"id"`
	Audio         string                    `json:"audio"`
	Level         string                    `json:"level"`
	Content       language.Multilingual     `json:"content"`
	Translated    string                    `json:"translated"`
	Vocabulary    string                    `json:"vocabulary"`
	CorrectAnswer string                    `json:"correctAnswer"`
	Options       []string                  `json:"options"`
	CorrectStreak int                       `json:"correctStreak"`
	IsFavorite    bool                      `json:"isFavorite"`
	IsMastered    bool                      `json:"isMastered"`
	UpdatedAt     *httprespond.TimeResponse `json:"updatedAt"`
	NextReviewAt  *httprespond.TimeResponse `json:"nextReviewAt"`
}

func (Exercise) FromDomain(exercise domain.Exercise) Exercise {
	return Exercise{
		ID:            exercise.ID,
		Audio:         exercise.Audio,
		Level:         exercise.Level,
		Content:       exercise.Content,
		Vocabulary:    exercise.Vocabulary,
		CorrectAnswer: exercise.CorrectAnswer,
		Options:       exercise.Options,
		CorrectStreak: exercise.CorrectStreak,
		IsFavorite:    exercise.IsFavorite,
		IsMastered:    exercise.IsMastered,
		UpdatedAt:     httprespond.NewTimeResponse(exercise.UpdatedAt),
		NextReviewAt:  httprespond.NewTimeResponse(exercise.NextReviewAt),
	}
}
