package command

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AnswerExerciseHandler struct {
	queueRepository domain.QueueRepository
	exerciseHub     domain.ExerciseHub
}

func NewAnswerExerciseHandler(queueRepository domain.QueueRepository, exerciseHub domain.ExerciseHub) AnswerExerciseHandler {
	return AnswerExerciseHandler{
		queueRepository: queueRepository,
		exerciseHub:     exerciseHub,
	}
}

func (h AnswerExerciseHandler) AnswerExercise(ctx *appcontext.AppContext, performerID, exerciseID string, req dto.AnswerExerciseRequest) (*dto.AnswerExerciseResponse, error) {
	ctx.Logger().Info("[command] new answer exercise request", appcontext.Fields{
		"performerID": performerID, "exerciseID": exerciseID,
		"isCorrect": req.IsCorrect, "completionTime": req.CompletionTime, "point": req.Point,
	})

	ctx.Logger().Text("prepare payload for calling English Hub")
	payload, err := domain.NewAnswerExercisePayload(performerID, exerciseID, req.IsCorrect)
	if err != nil {
		ctx.Logger().Error("failed to prepare payload for calling English Hub", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("call English Hub for answering")
	result, err := h.exerciseHub.AnswerExercise(ctx, *payload)
	if err != nil {
		ctx.Logger().Error("failed to call English Hub for answering", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("add queue task")
	if err = h.enqueueTasks(ctx, performerID, exerciseID, req.Point, req.CompletionTime); err != nil {
		ctx.Logger().Error("failed to add queue task", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done answer exercise request")
	return &dto.AnswerExerciseResponse{
		NextReviewAt: httprespond.NewTimeResponse(result.NextReviewAt),
	}, nil
}

func (h AnswerExerciseHandler) enqueueTasks(ctx *appcontext.AppContext, performerID, exerciseID string, point int64, completionTime int) error {
	ctx.Logger().Text("add task exerciseAnswered")
	if err := h.queueRepository.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPayload{
		UserID:         performerID,
		ExerciseID:     exerciseID,
		Point:          point,
		CompletionTime: completionTime,
	}); err != nil {
		ctx.Logger().Error("failed to add task exerciseAnswered", err, appcontext.Fields{})
	}

	return nil
}
