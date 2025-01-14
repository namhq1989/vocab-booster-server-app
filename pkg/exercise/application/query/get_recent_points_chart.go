package query

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/timezone"
)

type GetRecentPointsChartHandler struct {
	gamificationHub domain.GamificationHub
}

func NewGetRecentPointsChartHandler(gamificationHub domain.GamificationHub) GetRecentPointsChartHandler {
	return GetRecentPointsChartHandler{
		gamificationHub: gamificationHub,
	}
}

func (h GetRecentPointsChartHandler) GetRecentPointsChart(ctx *appcontext.AppContext, performerID string, tz timezone.Timezone, _ dto.GetRecentPointsChartRequest) (*dto.GetRecentPointsChartResponse, error) {
	ctx.Logger().Info("[query] new get recent points chart request", appcontext.Fields{"performerID": performerID, "timezone": tz.Identifier})

	ctx.Logger().Text("define time range")
	to := manipulation.Now(tz.Identifier)
	from := manipulation.StartOfDate(to.AddDate(0, 0, -6))

	ctx.Logger().Text("fetch points via grpc")
	points, err := h.gamificationHub.GetUserRecentPointsChart(ctx, performerID, tz.Identifier, from, to)
	if err != nil {
		ctx.Logger().Error("failed to get recent points chart", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.UserAggregatedPoint, 0)
	for _, point := range points {
		result = append(result, dto.UserAggregatedPoint{}.FromDomain(point))
	}

	ctx.Logger().Text("done get recent points chart request")
	return &dto.GetRecentPointsChartResponse{
		Points: result,
	}, nil
}
