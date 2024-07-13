package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		GetUserStats(ctx *appcontext.AppContext, req *gamificationpb.GetUserStatsRequest) (*gamificationpb.GetUserStatsResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		GetUserStatsHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userStatsRepository domain.UserStatsRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			GetUserStatsHandler: NewGetUserStatsHandler(userStatsRepository),
		},
	}
}
