package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type (
	Hubs interface {
		CreateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error)
		FindUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.FindUserSubscriptionRequest) (*subscriptionpb.FindUserSubscriptionResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		CreateUserSubscriptionHandler
		FindUserSubscriptionHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userSubscriptionHub domain.UserSubscriptionHub,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			CreateUserSubscriptionHandler: NewCreateUserSubscriptionHandler(userSubscriptionHub),
			FindUserSubscriptionHandler:   NewFindUserSubscriptionHandler(userSubscriptionHub),
		},
	}
}
