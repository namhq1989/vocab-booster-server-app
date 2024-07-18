package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
)

type UserAggregatedExerciseMapper struct{}

func (UserAggregatedExerciseMapper) FromGrpcToDomain(uae *exercisepb.UserAggregatedExercise) (*domain.UserAggregatedExercise, error) {
	result := domain.UserAggregatedExercise{
		Date:     uae.Date,
		Exercise: uae.Exercise,
	}

	return &result, nil
}
