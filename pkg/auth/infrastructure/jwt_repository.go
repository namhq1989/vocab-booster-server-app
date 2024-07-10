package infrastructure

import (
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
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
