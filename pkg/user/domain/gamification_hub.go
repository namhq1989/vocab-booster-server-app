package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type GamificationHub interface {
	GetUserStats(ctx *appcontext.AppContext, userID string) (*GamificationUserStats, error)
}

type GamificationUserStats struct {
	Point          int64
	CompletionTime int
}
