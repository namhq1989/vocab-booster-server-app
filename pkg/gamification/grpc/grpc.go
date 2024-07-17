package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"google.golang.org/grpc"
)

type server struct {
	hub App
	gamificationpb.UnimplementedGamificationServiceServer
}

var _ gamificationpb.GamificationServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, hub *Application) error {
	gamificationpb.RegisterGamificationServiceServer(registrar, server{hub: hub})
	return nil
}

func (s server) GetUserStats(bgCtx context.Context, req *gamificationpb.GetUserStatsRequest) (*gamificationpb.GetUserStatsResponse, error) {
	return s.hub.GetUserStats(appcontext.NewGRPC(bgCtx), req)
}

func (s server) GetUserRecentPointsChart(bgCtx context.Context, req *gamificationpb.GetUserRecentPointsChartRequest) (*gamificationpb.GetUserRecentPointsChartResponse, error) {
	return s.hub.GetUserRecentPointsChart(appcontext.NewGRPC(bgCtx), req)
}
