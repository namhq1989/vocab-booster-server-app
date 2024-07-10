package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type JourneyRepository interface {
	CreateJourney(ctx *appcontext.AppContext, journey Journey) error
	UpdateJourney(ctx *appcontext.AppContext, journey Journey) error
	FindUserCurrentJourney(ctx *appcontext.AppContext, userID string) (*Journey, error)
	FindJourneysByUserID(ctx *appcontext.AppContext, userID string) ([]Journey, error)
}

type Journey struct {
	ID         string
	UserID     string
	Lang       language.Language
	IsLearning bool
	Stats      JourneyStats
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type JourneyStats struct {
	Vocabulary             int
	SentenceComposed       int
	ExerciseMastered       int
	ExerciseCompletionTime int
	GainedPoints           int
}

func NewJourney(userID, lang string) (*Journey, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	dLang := language.ToLanguage(lang)
	if !dLang.IsValid() {
		return nil, apperrors.Common.InvalidLanguage
	}

	return &Journey{
		ID:         database.NewStringID(),
		UserID:     userID,
		Lang:       dLang,
		IsLearning: true,
		Stats:      JourneyStats{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (d *Journey) SetIsLearning(value bool) error {
	d.IsLearning = value
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Journey) SetStatsExerciseMastered(value int) error {
	if value < 0 {
		value = 0
	}
	d.Stats.ExerciseMastered = value
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Journey) IncreaseStatsVocabulary() error {
	d.Stats.Vocabulary++
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Journey) IncreaseStatsSentenceComposed() error {
	d.Stats.SentenceComposed++
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Journey) IncreaseStatsExerciseCompletionTime(value int) error {
	if value < 0 {
		value = 0
	}
	d.Stats.ExerciseCompletionTime += value
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Journey) IncreaseStatsGainedPoints(value int) error {
	if value < 0 {
		value = 0
	}
	d.Stats.GainedPoints += value
	d.UpdatedAt = time.Now()
	return nil
}
