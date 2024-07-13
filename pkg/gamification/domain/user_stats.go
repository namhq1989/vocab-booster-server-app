package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type UserStatsRepository interface {
	FindUserStats(ctx *appcontext.AppContext, userID string) (*UserStats, error)
	IncreaseUserStats(ctx *appcontext.AppContext, userID string, point int64, completionTime int) error
}

type UserStats struct {
	ID             string
	UserID         string
	Point          int64
	CompletionTime int
}
