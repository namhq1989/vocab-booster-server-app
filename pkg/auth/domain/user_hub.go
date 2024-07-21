package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserHub interface {
	FindUserByEmail(ctx *appcontext.AppContext, email string) (*User, error)
	CreateUser(ctx *appcontext.AppContext, name, email, timezone, providerSource, providerUid string) (string, error)
}

type User struct {
	ID   string
	Name string
}
