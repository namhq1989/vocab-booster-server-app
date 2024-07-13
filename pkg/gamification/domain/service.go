package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	ExerciseAnswered(ctx *appcontext.AppContext, point Point, completionTime CompletionTime) error
}
