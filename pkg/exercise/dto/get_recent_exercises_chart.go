package dto

import "github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"

type GetRecentExercisesChartRequest struct{}

type GetRecentExercisesChartResponse struct {
	Exercises []UserAggregatedExercise `json:"exercises"`
}

type UserAggregatedExercise struct {
	Date     string `json:"date"`
	Exercise int64  `json:"exercise"`
}

func (UserAggregatedExercise) FromDomain(uae domain.UserAggregatedExercise) UserAggregatedExercise {
	return UserAggregatedExercise{
		Date:     uae.Date,
		Exercise: uae.Exercise,
	}
}
