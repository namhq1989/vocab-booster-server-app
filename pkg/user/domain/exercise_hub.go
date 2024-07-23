package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type ExerciseHub interface {
	GetUserStats(ctx *appcontext.AppContext, userID, timezone string) (*ExerciseUserStats, error)
}

type ExerciseUserStats struct {
	Mastered         int
	WaitingForReview int
}
