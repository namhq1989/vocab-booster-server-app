package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"google.golang.org/grpc"
)

type server struct {
	hub App
	subscriptionpb.UnimplementedSubscriptionServiceServer
}

var _ subscriptionpb.SubscriptionServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, hub *Application) error {
	subscriptionpb.RegisterSubscriptionServiceServer(registrar, server{hub: hub})
	return nil
}

func (s server) CreateUserSubscription(bgCtx context.Context, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error) {
	return s.hub.CreateUserSubscription(appcontext.NewGRPC(bgCtx), req)
}

func (s server) GetUserSubscription(bgCtx context.Context, req *subscriptionpb.FindUserSubscriptionRequest) (*subscriptionpb.FindUserSubscriptionResponse, error) {
	return s.hub.FindUserSubscription(appcontext.NewGRPC(bgCtx), req)
}
