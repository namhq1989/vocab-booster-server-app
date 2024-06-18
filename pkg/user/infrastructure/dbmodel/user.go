package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (m User) ToDomain() domain.User {
	return domain.User{
		ID:        m.ID.Hex(),
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (User) FromDomain(user domain.User) (*User, error) {
	id, err := database.ObjectIDFromString(user.ID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
