package dbmodel

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Journey struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     primitive.ObjectID `bson:"userId"`
	Lang       language.Language  `bson:"lang"`
	IsLearning bool               `bson:"isLearning"`
	Stats      JourneyStats       `bson:"stats"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type JourneyStats struct {
	Vocabulary             int `bson:"vocabulary"`
	SentenceComposed       int `bson:"sentenceComposed"`
	ExerciseMastered       int `bson:"exerciseMastered"`
	ExerciseCompletionTime int `bson:"exerciseCompletionTime"`
	GainedPoints           int `bson:"gainedPoints"`
}

func (m Journey) ToDomain() domain.Journey {
	return domain.Journey{
		ID:         m.ID.Hex(),
		UserID:     m.UserID.Hex(),
		Lang:       m.Lang,
		IsLearning: m.IsLearning,
		Stats: domain.JourneyStats{
			Vocabulary:             m.Stats.Vocabulary,
			SentenceComposed:       m.Stats.SentenceComposed,
			ExerciseMastered:       m.Stats.ExerciseMastered,
			ExerciseCompletionTime: m.Stats.ExerciseCompletionTime,
			GainedPoints:           m.Stats.GainedPoints,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m Journey) FromDomain(journey domain.Journey) (*Journey, error) {
	id, err := database.ObjectIDFromString(journey.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	userID, err := database.ObjectIDFromString(journey.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &Journey{
		ID:         id,
		UserID:     userID,
		Lang:       journey.Lang,
		IsLearning: m.IsLearning,
		Stats: JourneyStats{
			Vocabulary:             journey.Stats.Vocabulary,
			SentenceComposed:       journey.Stats.SentenceComposed,
			ExerciseMastered:       journey.Stats.ExerciseMastered,
			ExerciseCompletionTime: journey.Stats.ExerciseCompletionTime,
			GainedPoints:           journey.Stats.GainedPoints,
		},
		CreatedAt: journey.CreatedAt,
		UpdatedAt: journey.UpdatedAt,
	}, nil
}
