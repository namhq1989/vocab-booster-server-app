package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
)

type ExerciseMapper struct{}

func (ExerciseMapper) FromGrpcToDomain(exercise *exercisepb.UserExercise) (*domain.Exercise, error) {
	result := domain.Exercise{
		ID:            exercise.GetId(),
		Audio:         exercise.GetAudio(),
		Level:         exercise.GetLevel(),
		Content:       dto.ConvertGrpcDataToMultilingual(exercise.GetContent()),
		Vocabulary:    exercise.GetVocabulary(),
		CorrectAnswer: exercise.GetCorrectAnswer(),
		Options:       exercise.GetOptions(),
		CorrectStreak: int(exercise.GetCorrectStreak()),
		IsFavorite:    exercise.GetIsFavorite(),
		IsMastered:    exercise.GetIsMastered(),
		UpdatedAt:     exercise.GetUpdatedAt().AsTime(),
		NextReviewAt:  exercise.GetNextReviewAt().AsTime(),
	}

	return &result, nil
}
