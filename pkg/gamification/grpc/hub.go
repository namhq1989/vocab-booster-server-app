package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		GetUserPoint(ctx *appcontext.AppContext, req *gamificationpb.GetUserPointRequest) (*gamificationpb.GetUserPointResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		GetUserPointHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userPointRepository domain.UserPointRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			GetUserPointHandler: NewGetUserPointHandler(userPointRepository),
		},
	}
}
