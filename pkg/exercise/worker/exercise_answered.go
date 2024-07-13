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
	if payload.Point > 0 {
		ctx.Logger().Text("add task addAnswerExercisePoint")
		if err := w.queueRepository.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
			UserID:     payload.UserID,
			ExerciseID: payload.ExerciseID,
			Point:      payload.Point,
		}); err != nil {
			ctx.Logger().Error("failed to add task addAnswerExercisePoint", err, appcontext.Fields{})
			return err
		}
	} else {
		ctx.Logger().Text("point is 0, skip add task addAnswerExercisePoint")
	}

	return nil
}
