package dto

type GetAccessTokenByUserIDRequest struct {
	UserID   string `query:"userId" validate:"required" message:"user_invalid_id"`
	Timezone string `query:"timezone"`
}

type GetAccessTokenByUserIDResponse struct {
	AccessToken string `json:"accessToken"`
}
