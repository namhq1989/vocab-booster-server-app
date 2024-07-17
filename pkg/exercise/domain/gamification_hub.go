package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GamificationHub interface {
	GetUserRecentPointsChart(ctx *appcontext.AppContext, userID string) ([]UserAggregatedPoint, error)
}

type UserAggregatedPoint struct {
	Date  string
	Point int64
}
