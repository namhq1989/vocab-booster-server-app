package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type (
	Commands interface {
		AnswerExercise(ctx *appcontext.AppContext, performerID, exerciseID string, req dto.AnswerExerciseRequest) (*dto.AnswerExerciseResponse, error)
	}
	Queries interface {
		GetExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetExercisesRequest) (*dto.GetExercisesResponse, error)
		GetReadyForReviewExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetReadyForReviewExercisesRequest) (*dto.GetReadyForReviewExercisesResponse, error)
		GetExerciseCollections(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetExerciseCollectionsRequest) (*dto.GetExerciseCollectionResponse, error)
		GetRecentPointsChart(ctx *appcontext.AppContext, performerID string, _ dto.GetRecentPointsChartRequest) (*dto.GetRecentPointsChartResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.AnswerExerciseHandler
	}
	appQueryHandler struct {
		query.GetExercisesHandler
		query.GetReadyForReviewExercisesHandler
		query.GetExerciseCollectionsHandler
		query.GetRecentPointsChartHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	queueRepository domain.QueueRepository,
	exerciseHub domain.ExerciseHub,
	gamificationHub domain.GamificationHub,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			AnswerExerciseHandler: command.NewAnswerExerciseHandler(queueRepository, exerciseHub),
		},
		appQueryHandler: appQueryHandler{
			GetExercisesHandler:               query.NewGetExercisesHandler(exerciseHub),
			GetReadyForReviewExercisesHandler: query.NewGetReadyForReviewExercisesHandler(exerciseHub),
			GetExerciseCollectionsHandler:     query.NewGetExerciseCollectionsHandler(exerciseHub),
			GetRecentPointsChartHandler:       query.NewGetRecentPointsChartHandler(gamificationHub),
		},
	}
}
