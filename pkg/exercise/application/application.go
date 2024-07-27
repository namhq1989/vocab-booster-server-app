package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"github.com/namhq1989/vocab-booster-utilities/timezone"
)

type (
	Commands interface {
		AnswerExercise(ctx *appcontext.AppContext, performerID, exerciseID string, tz timezone.Timezone, req dto.AnswerExerciseRequest) (*dto.AnswerExerciseResponse, error)
		ChangeExerciseFavorite(ctx *appcontext.AppContext, performerID, exerciseID string, req dto.ChangeExerciseFavoriteRequest) (*dto.ChangeExerciseFavoriteResponse, error)
	}
	Queries interface {
		GetExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetExercisesRequest) (*dto.GetExercisesResponse, error)
		GetReadyForReviewExercises(ctx *appcontext.AppContext, performerID string, lang language.Language, tz timezone.Timezone, _ dto.GetReadyForReviewExercisesRequest) (*dto.GetReadyForReviewExercisesResponse, error)
		GetExerciseCollections(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetExerciseCollectionsRequest) (*dto.GetExerciseCollectionResponse, error)
		GetRecentPointsChart(ctx *appcontext.AppContext, performerID string, tz timezone.Timezone, _ dto.GetRecentPointsChartRequest) (*dto.GetRecentPointsChartResponse, error)
		GetRecentExercisesChart(ctx *appcontext.AppContext, performerID string, tz timezone.Timezone, _ dto.GetRecentExercisesChartRequest) (*dto.GetRecentExercisesChartResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.AnswerExerciseHandler
		command.ChangeExerciseFavoriteHandler
	}
	appQueryHandler struct {
		query.GetExercisesHandler
		query.GetReadyForReviewExercisesHandler
		query.GetExerciseCollectionsHandler
		query.GetRecentPointsChartHandler
		query.GetRecentExercisesChartHandler
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
			AnswerExerciseHandler:         command.NewAnswerExerciseHandler(queueRepository, exerciseHub),
			ChangeExerciseFavoriteHandler: command.NewChangeExerciseFavoriteHandler(exerciseHub),
		},
		appQueryHandler: appQueryHandler{
			GetExercisesHandler:               query.NewGetExercisesHandler(exerciseHub),
			GetReadyForReviewExercisesHandler: query.NewGetReadyForReviewExercisesHandler(exerciseHub),
			GetExerciseCollectionsHandler:     query.NewGetExerciseCollectionsHandler(exerciseHub),
			GetRecentPointsChartHandler:       query.NewGetRecentPointsChartHandler(gamificationHub),
			GetRecentExercisesChartHandler:    query.NewGetRecentExercisesChartHandler(exerciseHub),
		},
	}
}
