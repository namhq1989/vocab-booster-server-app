package dto

type SignInWithGoogleRequest struct {
	Token    string `json:"token" validate:"required" message:"auth_invalid_google_token"`
	Timezone string `json:"timezone"`
}

type SignInWithGoogleResponse struct {
	AccessToken string `json:"accessToken"`
}
