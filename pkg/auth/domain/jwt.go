package domain

import "github.com/namhq1989/vocab-booster-server-app/core/appcontext"

type JwtRepository interface {
	GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error)
}
