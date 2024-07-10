package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Queries interface {
		GetExercises(ctx *appcontext.AppContext, performerID string, _ dto.GetExercisesRequest) (*dto.GetExercisesResponse, error)
	}
	Instance interface {
		Queries
	}

	appQueryHandler struct {
		query.GetExercisesHandler
	}
	Application struct {
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	exerciseHub domain.ExerciseHub,
) *Application {
	return &Application{
		appQueryHandler: appQueryHandler{
			GetExercisesHandler: query.NewGetExercisesHandler(exerciseHub),
		},
	}
}
