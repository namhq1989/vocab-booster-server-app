package dbmodel

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Avatar     string             `bson:"avatar"`
	Bio        string             `bson:"bio"`
	Visibility string             `bson:"visibility"`
	Providers  []UserProvider     `bson:"providers"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type UserProvider struct {
	Source string `bson:"source"`
	UID    string `bson:"uid"`
}

func (m User) ToDomain() domain.User {
	providers := make([]domain.UserProvider, 0)
	for _, provider := range m.Providers {
		providers = append(providers, domain.UserProvider{
			Source: provider.Source,
			UID:    provider.UID,
		})
	}

	return domain.User{
		ID:         m.ID.Hex(),
		Name:       m.Name,
		Email:      m.Email,
		Avatar:     m.Avatar,
		Bio:        m.Bio,
		Visibility: domain.Visibility(m.Visibility),
		Providers:  providers,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (User) FromDomain(user domain.User) (*User, error) {
	id, err := database.ObjectIDFromString(user.ID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	providers := make([]UserProvider, 0)
	for _, provider := range user.Providers {
		providers = append(providers, UserProvider{
			Source: provider.Source,
			UID:    provider.UID,
		})
	}

	return &User{
		ID:         id,
		Name:       user.Name,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Bio:        user.Bio,
		Visibility: user.Visibility.String(),
		Providers:  providers,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}
