package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseAnsweredHandler struct {
	pointRepository          domain.PointRepository
	completionTimeRepository domain.CompletionTimeRepository
	userStatsRepository      domain.UserStatsRepository
	service                  domain.Service
}

func NewExerciseAnsweredHandler(
	pointRepository domain.PointRepository,
	completionTimeRepository domain.CompletionTimeRepository,
	userStatsRepository domain.UserStatsRepository,
	service domain.Service,
) ExerciseAnsweredHandler {
	return ExerciseAnsweredHandler{
		pointRepository:          pointRepository,
		completionTimeRepository: completionTimeRepository,
		userStatsRepository:      userStatsRepository,
		service:                  service,
	}
}

func (w ExerciseAnsweredHandler) ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPoint) error {
	ctx.Logger().Text("new point model")
	point, err := domain.NewPoint(payload.UserID, payload.ExerciseID, "", domain.ActionAnswerExercise.String(), payload.Point)
	if err != nil {
		ctx.Logger().Error("failed to create new point model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("new completion time model")
	completionTime, err := domain.NewCompletionTime(payload.UserID, domain.ActionAnswerExercise.String(), payload.CompletionTime)
	if err != nil {
		ctx.Logger().Error("failed to create new completion time model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist in db with transaction")
	if err = w.service.ExerciseAnswered(ctx, *point, *completionTime); err != nil {
		ctx.Logger().Error("failed to persist in db with transaction", err, appcontext.Fields{})
		return err
	}

	return nil
}
