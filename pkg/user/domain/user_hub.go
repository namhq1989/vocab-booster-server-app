package domain

import (
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
)

type UserHub interface {
	FindUserByEmail(ctx *appcontext.AppContext, email string) (*User, error)
	FindUserByID(ctx *appcontext.AppContext, userID string) (*User, error)
	CreateUser(ctx *appcontext.AppContext, user User) error
}
