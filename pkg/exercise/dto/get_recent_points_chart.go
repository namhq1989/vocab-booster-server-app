package dto

import "github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"

type GetRecentPointsChartRequest struct{}

type GetRecentPointsChartResponse struct {
	Points []UserAggregatedPoint `json:"points"`
}

type UserAggregatedPoint struct {
	Date  string `json:"date"`
	Point int64  `json:"point"`
}

func (UserAggregatedPoint) FromDomain(uap domain.UserAggregatedPoint) UserAggregatedPoint {
	return UserAggregatedPoint{
		Date:  uap.Date,
		Point: uap.Point,
	}
}
