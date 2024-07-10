package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserRepository interface {
	FindUserByID(ctx *appcontext.AppContext, userID string) (*User, error)
	UpdateUser(ctx *appcontext.AppContext, user User) error
}

type User struct {
	ID         string
	Name       string
	Email      string
	Bio        string
	Visibility Visibility
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUser(name, email string) (*User, error) {
	if !validation.IsValidUserName(name) {
		return nil, apperrors.Common.InvalidName
	}

	if !validation.IsValidEmail(email) {
		return nil, apperrors.Common.InvalidEmail
	}

	return &User{
		ID:         database.NewStringID(),
		Name:       name,
		Email:      email,
		Visibility: VisibilityPublic,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (d *User) SetName(name string) error {
	if !validation.IsValidUserName(name) {
		return apperrors.Common.InvalidName
	}

	d.Name = name
	d.UpdatedAt = time.Now()
	return nil
}

func (d *User) SetBio(bio string) error {
	d.Bio = bio
	d.UpdatedAt = time.Now()
	return nil
}

func (d *User) SetVisibility(visibility string) error {
	dVisibility := ToVisibility(visibility)
	if !dVisibility.IsValid() {
		return apperrors.User.InvalidVisibility
	}

	d.Visibility = dVisibility
	d.UpdatedAt = time.Now()
	return nil
}

func (d *User) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}
