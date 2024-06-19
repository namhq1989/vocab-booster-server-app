package application

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/dto"
)

type (
	Commands interface {
	}
	Queries interface {
		GetSubscriptionPlans(ctx *appcontext.AppContext, performerID string, _ dto.GetSubscriptionPlansRequest) (*dto.GetSubscriptionPlansResponse, error)
	}
	App interface {
		Queries
	}

	appCommandHandlers struct {
	}
	appQueryHandler struct {
		query.GetSubscriptionPlansHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ App = (*Application)(nil)

func New() *Application {
	return &Application{
		appQueryHandler: appQueryHandler{
			GetSubscriptionPlansHandler: query.NewGetSubscriptionPlansHandler(),
		},
	}
}
