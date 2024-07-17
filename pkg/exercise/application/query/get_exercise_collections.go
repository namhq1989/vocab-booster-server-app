package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetExerciseCollectionsHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewGetExerciseCollectionsHandler(exerciseHub domain.ExerciseHub) GetExerciseCollectionsHandler {
	return GetExerciseCollectionsHandler{
		exerciseHub: exerciseHub,
	}
}

func (h GetExerciseCollectionsHandler) GetExerciseCollections(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetExerciseCollectionsRequest) (*dto.GetExerciseCollectionResponse, error) {
	ctx.Logger().Info("[query] new get exercise collections request", appcontext.Fields{"performerID": performerID, "lang": lang.String()})

	ctx.Logger().Text("fetch exercise collections via grpc")
	collections, err := h.exerciseHub.GetExerciseCollections(ctx, performerID, lang.String())
	if err != nil {
		ctx.Logger().Error("failed to get exercise collections", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response data")
	result := make([]dto.ExerciseCollection, 0)
	for _, collection := range collections {
		result = append(result, dto.ExerciseCollection{}.FromDomain(collection))
	}

	ctx.Logger().Text("done get exercise collections request")
	return &dto.GetExerciseCollectionResponse{
		Collections: result,
	}, nil
}
