package dto

import "github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Bio        string `json:"bio"`
	Visibility string `json:"visibility"`
}

func (User) FromDomain(user domain.User) User {
	return User{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Bio:        user.Bio,
		Visibility: user.Visibility.String(),
	}
}