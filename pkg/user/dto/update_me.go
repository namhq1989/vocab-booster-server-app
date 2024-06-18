package dto

type UpdateMeRequest struct {
	Name       string `json:"name" validate:"required" message:"invalid_name"`
	Bio        string `json:"bio"`
	Visibility string `json:"visibility" validate:"required" message:"user_invalid_visibility"`
}

type UpdateMeResponse struct{}
