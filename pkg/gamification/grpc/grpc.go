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

func (s server) GetUserPoint(bgCtx context.Context, req *gamificationpb.GetUserPointRequest) (*gamificationpb.GetUserPointResponse, error) {
	return s.hub.GetUserPoint(appcontext.NewGRPC(bgCtx), req)
}
