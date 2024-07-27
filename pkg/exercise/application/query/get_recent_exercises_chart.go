package query

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/timezone"
)

type GetRecentExercisesChartHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewGetRecentExercisesChartHandler(exerciseHub domain.ExerciseHub) GetRecentExercisesChartHandler {
	return GetRecentExercisesChartHandler{
		exerciseHub: exerciseHub,
	}
}

func (h GetRecentExercisesChartHandler) GetRecentExercisesChart(ctx *appcontext.AppContext, performerID string, tz timezone.Timezone, _ dto.GetRecentExercisesChartRequest) (*dto.GetRecentExercisesChartResponse, error) {
	ctx.Logger().Info("[query] new get recent exercises chart request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("define time range")
	to := manipulation.Now(tz.Identifier)
	from := manipulation.StartOfDate(to.AddDate(0, 0, -6))

	ctx.Logger().Text("fetch exercises via grpc")
	exercises, err := h.exerciseHub.AggregateUserExercisesInTimeRange(ctx, performerID, tz.Identifier, from, to)
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
