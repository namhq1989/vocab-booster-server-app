package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type QueueRepository interface {
	ExerciseAnswered(ctx *appcontext.AppContext, payload QueueExerciseAnsweredPayload) error
	AddAnswerExercisePoint(ctx *appcontext.AppContext, payload QueueAddAnswerExercisePoint) error
}

type QueueExerciseAnsweredPayload struct {
	UserID         string
	ExerciseID     string
	Point          int64
	CompletionTime int
}

type QueueAddAnswerExercisePoint struct {
	UserID     string
	ExerciseID string
	Point      int64
}
