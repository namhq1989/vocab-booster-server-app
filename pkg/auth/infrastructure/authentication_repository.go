package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/authentication"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AuthenticationRepository struct {
	authentication authentication.Operations
}

func NewAuthenticationRepository(authentication *authentication.Authentication) AuthenticationRepository {
	return AuthenticationRepository{
		authentication: authentication,
	}
}

func (r AuthenticationRepository) GetUserInfoWithToken(ctx *appcontext.AppContext, token string) (*domain.AuthenticationUser, error) {
	user, err := r.authentication.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &domain.AuthenticationUser{
		UID:   user.UID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
