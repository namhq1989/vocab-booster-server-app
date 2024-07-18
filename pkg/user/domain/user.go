package domain

import (
	"fmt"
	"slices"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"

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
	Avatar     string
	Visibility Visibility
	Providers  []UserProvider
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserProvider struct {
	Source string
	UID    string
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
		Avatar:     randomAvatar(),
		Visibility: VisibilityPublic,
		CreatedAt:  manipulation.Now(),
		UpdatedAt:  manipulation.Now(),
	}, nil
}

func (d *User) SetName(name string) error {
	if !validation.IsValidUserName(name) {
		return apperrors.Common.InvalidName
	}

	d.Name = name
	d.UpdatedAt = manipulation.Now()
	return nil
}

func (d *User) SetBio(bio string) error {
	d.Bio = bio
	d.UpdatedAt = manipulation.Now()
	return nil
}

func (d *User) SetVisibility(visibility string) error {
	dVisibility := ToVisibility(visibility)
	if !dVisibility.IsValid() {
		return apperrors.User.InvalidVisibility
	}

	d.Visibility = dVisibility
	d.UpdatedAt = manipulation.Now()
	return nil
}

func (d *User) SetProvider(source, uid string) error {
	index := slices.IndexFunc(d.Providers, func(provider UserProvider) bool { return provider.Source == source })
	if index >= 0 {
		d.Providers[index].UID = uid
	} else {
		d.Providers = append(d.Providers, UserProvider{
			Source: source,
			UID:    uid,
		})
	}
	d.UpdatedAt = manipulation.Now()
	return nil
}

const totalAvatars = 26

func randomAvatar() string {
	r := manipulation.RandomIntInRange(1, totalAvatars)
	return fmt.Sprintf("%d", r)
}

func (d *User) SetAvatar(value string) {
	d.Avatar = value
	d.UpdatedAt = manipulation.Now()
}

func (d *User) SetUpdatedAt() {
	d.UpdatedAt = manipulation.Now()
}
