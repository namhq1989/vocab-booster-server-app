package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		GetUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.GetUserSubscriptionRequest) (*subscriptionpb.GetUserSubscriptionResponse, error)

		CreateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error)
		UpdateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.UpdateUserSubscriptionRequest) (*subscriptionpb.UpdateUserSubscriptionResponse, error)
		CanPerformAction(ctx *appcontext.AppContext, req *subscriptionpb.CanPerformActionRequest) (*subscriptionpb.CanPerformActionResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		GetUserSubscriptionHandler

		CreateUserSubscriptionHandler
		UpdateUserSubscriptionHandler
		CanPerformActionHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userSubscriptionRepository domain.UserSubscriptionRepository,
	userSubscriptionHistoryRepository domain.UserSubscriptionHistoryRepository,
	userSubscriptionHub domain.UserSubscriptionHub,
	service domain.Service,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			GetUserSubscriptionHandler: NewGetUserSubscriptionHandler(userSubscriptionHub),

			CreateUserSubscriptionHandler: NewCreateUserSubscriptionHandler(userSubscriptionHub),
			UpdateUserSubscriptionHandler: NewUpdateUserSubscriptionHandler(userSubscriptionRepository, userSubscriptionHistoryRepository),
			CanPerformActionHandler:       NewCanPerformActionHandler(service),
		},
	}
}
