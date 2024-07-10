package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Commands interface {
		SignInWithGoogle(ctx *appcontext.AppContext, req dto.SignInWithGoogleRequest) (*dto.SignInWithGoogleResponse, error)
	}
	Queries interface {
		GetAccessTokenByUserID(ctx *appcontext.AppContext, req dto.GetAccessTokenByUserIDRequest) (*dto.GetAccessTokenByUserIDResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.SignInWithGoogleHandler
	}
	appQueryHandler struct {
		query.GetAccessTokenByUserIDHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	ssoRepository domain.SSORepository,
	jwtRepository domain.JwtRepository,
	userHub domain.UserHub,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			SignInWithGoogleHandler: command.NewSignInWithGoogleHandler(ssoRepository, jwtRepository, userHub),
		},
		appQueryHandler: appQueryHandler{
			GetAccessTokenByUserIDHandler: query.NewGetAccessTokenByUserIDHandler(jwtRepository),
		},
	}
}
