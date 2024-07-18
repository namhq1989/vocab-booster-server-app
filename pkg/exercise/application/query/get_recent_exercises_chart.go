package query

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetRecentExercisesChartHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewGetRecentExercisesChartHandler(exerciseHub domain.ExerciseHub) GetRecentExercisesChartHandler {
	return GetRecentExercisesChartHandler{
		exerciseHub: exerciseHub,
	}
}

func (h GetRecentExercisesChartHandler) GetRecentExercisesChart(ctx *appcontext.AppContext, performerID string, _ dto.GetRecentExercisesChartRequest) (*dto.GetRecentExercisesChartResponse, error) {
	ctx.Logger().Info("[query] new get recent exercises chart request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("define time range")
	to := manipulation.Now()
	from := manipulation.StartOfDate(to.AddDate(0, 0, -6))

	ctx.Logger().Text("fetch exercises via grpc")
	exercises, err := h.exerciseHub.AggregateUserExercisesInTimeRange(ctx, performerID, from, to)
	if err != nil {
		ctx.Logger().Error("failed to get recent exercises chart", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.UserAggregatedExercise, 0)
	for _, exercise := range exercises {
		result = append(result, dto.UserAggregatedExercise{}.FromDomain(exercise))
	}

	ctx.Logger().Text("done get recent exercises chart request")
	return &dto.GetRecentExercisesChartResponse{
		Exercises: result,
	}, nil
}
