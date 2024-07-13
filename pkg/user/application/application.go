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
		ChangeAvatar(ctx *appcontext.AppContext, performerID string, req dto.ChangeAvatarRequest) (*dto.ChangeAvatarResponse, error)
	}
	Queries interface {
		GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error)
		GetStats(ctx *appcontext.AppContext, performerID string, _ dto.GetStatsRequest) (*dto.GetStatsResponse, error)
	}
	App interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.UpdateMeHandler
		command.ChangeAvatarHandler
	}
	appQueryHandler struct {
		query.GetMeHandler
		query.GetStatsHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	gamificationHub domain.GamificationHub,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			UpdateMeHandler:     command.NewUpdateMeHandler(userRepository),
			ChangeAvatarHandler: command.NewChangeAvatarHandler(userRepository),
		},
		appQueryHandler: appQueryHandler{
			GetMeHandler:    query.NewGetMeHandler(userRepository),
			GetStatsHandler: query.NewGetStatsHandler(gamificationHub),
		},
	}
}
