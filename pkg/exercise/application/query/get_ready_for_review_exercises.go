package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetReadyForReviewExercisesHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewGetReadyForReviewExercisesHandler(exerciseHub domain.ExerciseHub) GetReadyForReviewExercisesHandler {
	return GetReadyForReviewExercisesHandler{
		exerciseHub: exerciseHub,
	}
}

func (h GetReadyForReviewExercisesHandler) GetReadyForReviewExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetReadyForReviewExercisesRequest) (*dto.GetReadyForReviewExercisesResponse, error) {
	ctx.Logger().Info("[query] new get ready for review exercises request", appcontext.Fields{"performerID": performerID, "lang": lang.String()})
	exercises, err := h.exerciseHub.GetReadyForReviewExercises(ctx, performerID, lang.String())
	if err != nil {
		ctx.Logger().Error("failed to get ready for review exercises", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.Exercise, 0)
	for _, exercise := range exercises {
		result = append(result, dto.Exercise{}.FromDomain(exercise))
	}

	ctx.Logger().Text("done get ready for review exercises request")
	return &dto.GetReadyForReviewExercisesResponse{
		Exercises: result,
	}, nil
}
