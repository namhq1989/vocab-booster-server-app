package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/userpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserHub struct {
	client userpb.UserServiceClient
}

func NewUserHub(client userpb.UserServiceClient) UserHub {
	return UserHub{
		client: client,
	}
}

func (r UserHub) FindUserByEmail(ctx *appcontext.AppContext, email string) (*domain.User, error) {
	resp, err := r.client.FindUserByEmail(ctx.Context(), &userpb.FindUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	user := resp.GetUser()
	if user == nil {
		return nil, nil
	}
	return &domain.User{
		ID:   user.GetId(),
		Name: user.GetName(),
	}, nil
}

func (r UserHub) CreateUser(ctx *appcontext.AppContext, name, email, providerSource, providerUid string) (string, error) {
	resp, err := r.client.CreateUser(ctx.Context(), &userpb.CreateUserRequest{
		Name:           name,
		Email:          email,
		ProviderSource: providerSource,
		ProviderUid:    providerUid,
	})
	if err != nil {
		return "", err
	}
	return resp.GetId(), nil
}
