package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
)

type ExerciseCollectionMapper struct{}

func (ExerciseCollectionMapper) FromGrpcToDomain(collection *exercisepb.ExerciseCollection) (*domain.ExerciseCollection, error) {
	result := domain.ExerciseCollection{
		ID:              collection.GetId(),
		Name:            collection.GetName(),
		Slug:            collection.GetSlug(),
		Translated:      collection.GetTranslated(),
		StatsExercises:  int(collection.GetStatsExercises()),
		StatsInteracted: int(collection.GetStatsInteracted()),
		Image:           collection.GetImage(),
	}

	return &result, nil
}
