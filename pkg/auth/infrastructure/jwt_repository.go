package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
)

type JwtRepository struct {
	jwt appjwt.Operations
}

func NewJwtRepository(jwt appjwt.Operations) JwtRepository {
	return JwtRepository{
		jwt: jwt,
	}
}

func (r JwtRepository) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	return r.jwt.GenerateAccessToken(ctx, userID)
}
