package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetExercisesHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewGetExercisesHandler(exerciseHub domain.ExerciseHub) GetExercisesHandler {
	return GetExercisesHandler{
		exerciseHub: exerciseHub,
	}
}

func (h GetExercisesHandler) GetExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.GetExercisesRequest) (*dto.GetExercisesResponse, error) {
	ctx.Logger().Info("[query] new get exercises request", appcontext.Fields{"performerID": performerID, "lang": lang.String(), "level": req.Level})

	ctx.Logger().Text("fetch exercises via grpc")
	exercises, err := h.exerciseHub.GetExercises(ctx, performerID, lang.String(), req.Level)
	if err != nil {
		ctx.Logger().Error("failed to get exercises", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.Exercise, 0)
	for _, exercise := range exercises {
		result = append(result, dto.Exercise{}.FromDomain(exercise))
	}

	ctx.Logger().Text("done get exercises request")
	return &dto.GetExercisesResponse{
		Exercises: result,
	}, nil
}
