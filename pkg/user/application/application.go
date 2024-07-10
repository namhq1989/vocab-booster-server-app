package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Commands interface {
		UpdateMe(ctx *appcontext.AppContext, performerID string, req dto.UpdateMeRequest) (*dto.UpdateMeResponse, error)

		CreateJourney(ctx *appcontext.AppContext, performerID string, req dto.CreateJourneyRequest) (*dto.CreateJourneyResponse, error)
		SwitchJourney(ctx *appcontext.AppContext, performerID string, req dto.SwitchJourneyRequest) (*dto.SwitchJourneyResponse, error)
		GetJourneys(ctx *appcontext.AppContext, performerID string, _ dto.GetJourneysRequest) (*dto.GetJourneysResponse, error)
	}
	Queries interface {
		GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error)
	}
	App interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.UpdateMeHandler

		command.CreateJourneyHandler
		command.SwitchJourneyHandler
	}
	appQueryHandler struct {
		query.GetMeHandler

		query.GetJourneysHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	journeyRepository domain.JourneyRepository,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			UpdateMeHandler: command.NewUpdateMeHandler(userRepository),

			CreateJourneyHandler: command.NewCreateJourneyHandler(journeyRepository),
			SwitchJourneyHandler: command.NewSwitchJourneyHandler(journeyRepository),
		},
		appQueryHandler: appQueryHandler{
			GetMeHandler: query.NewGetMeHandler(userRepository),

			GetJourneysHandler: query.NewGetJourneysHandler(journeyRepository),
		},
	}
}
