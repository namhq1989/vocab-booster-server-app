package command

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ChangeExerciseFavoriteHandler struct {
	exerciseHub domain.ExerciseHub
}

func NewChangeExerciseFavoriteHandler(exerciseHub domain.ExerciseHub) ChangeExerciseFavoriteHandler {
	return ChangeExerciseFavoriteHandler{
		exerciseHub: exerciseHub,
	}
}

func (h ChangeExerciseFavoriteHandler) ChangeExerciseFavorite(ctx *appcontext.AppContext, performerID, exerciseID string, req dto.ChangeExerciseFavoriteRequest) (*dto.ChangeExerciseFavoriteResponse, error) {
	ctx.Logger().Info("[command] new change exercise favorite request", appcontext.Fields{"userID": performerID, "exerciseID": exerciseID, "isFavorite": req.IsFavorite})

	ctx.Logger().Text("call English hub for changing isFavorite")
	isFavorite, err := h.exerciseHub.ChangeExerciseFavorite(ctx, performerID, exerciseID, req.IsFavorite)
	if err != nil {
		ctx.Logger().Error("failed to call English hub for changing isFavorite", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done change exercise favorite request")
	return &dto.ChangeExerciseFavoriteResponse{
		IsFavorite: isFavorite,
	}, nil
}
