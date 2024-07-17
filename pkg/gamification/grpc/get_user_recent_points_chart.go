package grpc

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserRecentPointsChartHandler struct {
	pointRepository domain.PointRepository
}

func NewGetUserRecentPointsChartHandler(pointRepository domain.PointRepository) GetUserRecentPointsChartHandler {
	return GetUserRecentPointsChartHandler{
		pointRepository: pointRepository,
	}
}

func (h GetUserRecentPointsChartHandler) GetUserRecentPointsChart(ctx *appcontext.AppContext, req *gamificationpb.GetUserRecentPointsChartRequest) (*gamificationpb.GetUserRecentPointsChartResponse, error) {
	ctx.Logger().Info("[hub] new get user recent points chart request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("define time range")
	to := time.Now()
	from := manipulation.StartOfDate(to.AddDate(0, 0, -6))

	ctx.Logger().Text("find in db")
	points, err := h.pointRepository.AggregateUserPointsInTimeRange(ctx, req.GetUserId(), from, to)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response data")
	var result = make([]*gamificationpb.UserAggregatedPoint, 0)
	for _, point := range points {
		result = append(result, &gamificationpb.UserAggregatedPoint{
			Date:  point.Date,
			Point: point.Point,
		})
	}

	ctx.Logger().Text("done get user recent points chart request")
	return &gamificationpb.GetUserRecentPointsChartResponse{
		Points: result,
	}, nil
}
