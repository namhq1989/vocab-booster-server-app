package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetRecentPointsChartHandler struct {
	gamificationHub domain.GamificationHub
}

func NewGetRecentPointsChartHandler(gamificationHub domain.GamificationHub) GetRecentPointsChartHandler {
	return GetRecentPointsChartHandler{
		gamificationHub: gamificationHub,
	}
}

func (h GetRecentPointsChartHandler) GetRecentPointsChart(ctx *appcontext.AppContext, performerID string, _ dto.GetRecentPointsChartRequest) (*dto.GetRecentPointsChartResponse, error) {
	ctx.Logger().Info("[query] new get recent points chart request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("fetch points via grpc")
	points, err := h.gamificationHub.GetUserRecentPointsChart(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get recent points chart", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.UserAggregatedPoint, 0)
	for _, point := range points {
		result = append(result, dto.UserAggregatedPoint{}.FromDomain(point))
	}

	ctx.Logger().Text("done get ready for review exercises request")
	return &dto.GetRecentPointsChartResponse{
		Points: result,
	}, nil
}
