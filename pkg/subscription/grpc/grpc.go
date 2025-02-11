package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
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

func (s server) GetUserSubscription(bgCtx context.Context, req *subscriptionpb.GetUserSubscriptionRequest) (*subscriptionpb.GetUserSubscriptionResponse, error) {
	return s.hub.GetUserSubscription(appcontext.NewGRPC(bgCtx), req)
}

func (s server) CreateUserSubscription(bgCtx context.Context, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error) {
	return s.hub.CreateUserSubscription(appcontext.NewGRPC(bgCtx), req)
}

func (s server) UpdateUserSubscription(bgCtx context.Context, req *subscriptionpb.UpdateUserSubscriptionRequest) (*subscriptionpb.UpdateUserSubscriptionResponse, error) {
	return s.hub.UpdateUserSubscription(appcontext.NewGRPC(bgCtx), req)
}
