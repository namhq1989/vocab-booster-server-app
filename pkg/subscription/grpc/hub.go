package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
)

type (
	Hubs interface {
		FindUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.FindUserSubscriptionRequest) (*subscriptionpb.FindUserSubscriptionResponse, error)

		CreateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.CreateUserSubscriptionRequest) (*subscriptionpb.CreateUserSubscriptionResponse, error)
		UpdateUserSubscription(ctx *appcontext.AppContext, req *subscriptionpb.UpdateUserSubscriptionRequest) (*subscriptionpb.UpdateUserSubscriptionResponse, error)
		CanPerformAction(ctx *appcontext.AppContext, req *subscriptionpb.CanPerformActionRequest) (*subscriptionpb.CanPerformActionResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		FindUserSubscriptionHandler

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
			FindUserSubscriptionHandler: NewFindUserSubscriptionHandler(userSubscriptionHub),

			CreateUserSubscriptionHandler: NewCreateUserSubscriptionHandler(userSubscriptionHub),
			UpdateUserSubscriptionHandler: NewUpdateUserSubscriptionHandler(userSubscriptionRepository, userSubscriptionHistoryRepository),
			CanPerformActionHandler:       NewCanPerformActionHandler(service),
		},
	}
}
