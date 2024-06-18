package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/userpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
)

type CreateUserHandler struct {
	userHub domain.UserHub
}

func NewCreateUserHandler(userHub domain.UserHub) CreateUserHandler {
	return CreateUserHandler{
		userHub: userHub,
	}
}

func (h CreateUserHandler) CreateUser(ctx *appcontext.AppContext, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx.Logger().Info("[hub] new create user request", appcontext.Fields{"name": req.GetName(), "email": req.GetEmail()})

	ctx.Logger().Text("create new user's model")
	user, err := domain.NewUser(req.GetName(), req.GetEmail())
	if err != nil {
		ctx.Logger().Error("failed to create new user's model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist user in db")
	err = h.userHub.CreateUser(ctx, *user)
	if err != nil {
		ctx.Logger().Error("failed to persist user in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("enqueue tasks")
	if err = h.enqueueTasks(ctx, *user); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Info("done create user request", appcontext.Fields{"userID": user.ID})
	return &userpb.CreateUserResponse{
		Id: user.ID,
	}, nil
}

func (h CreateUserHandler) enqueueTasks(ctx *appcontext.AppContext, user domain.User) error {
	// DO SOMETHING WITH QUEUE
	return nil
}
