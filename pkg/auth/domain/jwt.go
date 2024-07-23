package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type JwtRepository interface {
	GenerateAccessToken(ctx *appcontext.AppContext, userID, timezone string) (string, error)
}
