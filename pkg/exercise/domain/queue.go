package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type QueueRepository interface {
	ExerciseAnswered(ctx *appcontext.AppContext, payload QueueExerciseAnsweredPayload) error
}

type QueueExerciseAnsweredPayload struct {
	UserID         string
	Points         int
	CompletionTime int
}
