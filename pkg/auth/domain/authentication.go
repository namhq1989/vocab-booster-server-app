package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AuthenticationRepository interface {
	GetUserInfoWithToken(ctx *appcontext.AppContext, token string) (*AuthenticationUser, error)
}

type AuthenticationUser struct {
	UID            string
	Email          string
	Name           string
	ProviderSource string
	ProviderUID    string
}
