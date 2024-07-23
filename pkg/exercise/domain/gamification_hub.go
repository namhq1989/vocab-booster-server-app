package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GamificationHub interface {
	GetUserRecentPointsChart(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]UserAggregatedPoint, error)
}

type UserAggregatedPoint struct {
	Date  string
	Point int64
}
