package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/userpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		FindUserByEmail(ctx *appcontext.AppContext, req *userpb.FindUserByEmailRequest) (*userpb.FindUserByEmailResponse, error)
		FindUserByID(ctx *appcontext.AppContext, req *userpb.FindUserByIDRequest) (*userpb.FindUserByIDResponse, error)
		CreateUser(ctx *appcontext.AppContext, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		FindUserByIDHandler
		FindUserByEmailHandler
		CreateUserHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	userHub domain.UserHub,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			FindUserByIDHandler:    NewFindUserByIDHandler(userHub),
			FindUserByEmailHandler: NewFindUserByEmailHandler(userHub),
			CreateUserHandler:      NewCreateUserHandler(userHub),
		},
	}
}
