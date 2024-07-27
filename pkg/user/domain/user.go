package domain

import (
	"fmt"
	"slices"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/timezone"
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
	Timezone   timezone.Timezone
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserProvider struct {
	Source string
	UID    string
}

func NewUser(name, email, tz string) (*User, error) {
	if !validation.IsValidUserName(name) {
		return nil, apperrors.Common.InvalidName
	}

	if !validation.IsValidEmail(email) {
		return nil, apperrors.Common.InvalidEmail
	}

	dTimezone, _ := timezone.GetTimezoneData(tz)

	return &User{
		ID:         database.NewStringID(),
		Name:       name,
		Email:      email,
		Avatar:     randomAvatar(),
		Visibility: VisibilityPublic,
		Timezone:   *dTimezone,
		CreatedAt:  manipulation.NowUTC(),
		UpdatedAt:  manipulation.NowUTC(),
	}, nil
}

func (d *User) SetName(name string) error {
	if !validation.IsValidUserName(name) {
		return apperrors.Common.InvalidName
	}

	d.Name = name
	d.SetUpdatedAt()
	return nil
}

func (d *User) SetBio(bio string) error {
	d.Bio = bio
	d.SetUpdatedAt()
	return nil
}

func (d *User) SetVisibility(visibility string) error {
	dVisibility := ToVisibility(visibility)
	if !dVisibility.IsValid() {
		return apperrors.User.InvalidVisibility
	}

	d.Visibility = dVisibility
	d.SetUpdatedAt()
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
	d.SetUpdatedAt()
	return nil
}

const totalAvatars = 26

func randomAvatar() string {
	r := manipulation.RandomIntInRange(1, totalAvatars)
	return fmt.Sprintf("%d", r)
}

func (d *User) SetAvatar(value string) {
	d.Avatar = value
	d.SetUpdatedAt()
}

func (d *User) SetTimezone(tz string) {
	dTimezone, _ := timezone.GetTimezoneData(tz)

	d.Timezone = *dTimezone
	d.SetUpdatedAt()
}

func (d *User) SetUpdatedAt() {
	d.UpdatedAt = manipulation.NowUTC()
}
