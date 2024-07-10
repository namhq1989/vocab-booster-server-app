package dto

type SwitchJourneyRequest struct {
	Lang string `json:"lang" validate:"required" message:"invalid_lang"`
}

type SwitchJourneyResponse struct {
	ID string `json:"id"`
}
