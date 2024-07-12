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

func (ExerciseAnsweredHandler) ExerciseAnswered(ctx *appcontext.AppContext, payload domain.QueueExerciseAnsweredPayload) error {
	ctx.Logger().Text("** DO NOTHING **")
	return nil
}
