package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseAnsweredHandler struct {
	queueRepository domain.QueueRepository
}

func NewExerciseAnsweredHandler(
	queueRepository domain.QueueRepository,
) ExerciseAnsweredHandler {
	return ExerciseAnsweredHandler{
		queueRepository: queueRepository,
	}
}

func (w ExerciseAnsweredHandler) ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPayload) error {
	ctx.Logger().Text("add task gamification.exerciseAnswered")
	if err := w.queueRepository.GamificationExerciseAnswered(ctx, domain.QueueExerciseAnsweredPayload{
		UserID:         payload.UserID,
		ExerciseID:     payload.ExerciseID,
		Point:          payload.Point,
		CompletionTime: payload.CompletionTime,
	}); err != nil {
		ctx.Logger().Error("failed to add task gamification.exerciseAnswered", err, appcontext.Fields{})
		return err
	}

	return nil
}
