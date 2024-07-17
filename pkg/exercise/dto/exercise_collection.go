package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
)

type ExerciseCollection struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Translated      string `json:"translated"`
	StatsExercises  int    `json:"statsExercises"`
	StatsInteracted int    `json:"statsInteracted"`
	Image           string `json:"image"`
}

func (ExerciseCollection) FromDomain(collection domain.ExerciseCollection) ExerciseCollection {
	return ExerciseCollection{
		ID:              collection.ID,
		Name:            collection.Name,
		Slug:            collection.Slug,
		Translated:      collection.Translated,
		StatsExercises:  collection.StatsExercises,
		StatsInteracted: collection.StatsInteracted,
		Image:           collection.Image,
	}
}
