package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
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
	ctx.Logger().Info("[hub] new get user recent points chart request", appcontext.Fields{"userID": req.GetUserId(), "from": req.GetFrom().AsTime(), "to": req.GetTo().AsTime()})

	ctx.Logger().Text("find in db")
	points, err := h.pointRepository.AggregateUserPointsInTimeRange(ctx, req.GetUserId(), req.GetFrom().AsTime(), req.GetTo().AsTime())
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
