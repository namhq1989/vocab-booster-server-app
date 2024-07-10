package dto

type CreateJourneyRequest struct {
	Lang string `json:"lang" validate:"required" message:"invalid_lang"`
}

type CreateJourneyResponse struct {
	ID string `json:"id"`
}
