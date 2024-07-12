package dto

type ChangeAvatarRequest struct {
	Avatar string `json:"avatar" validate:"required" message:"invalid_avatar"`
}

type ChangeAvatarResponse struct{}
