package dto

type GetJourneysRequest struct{}

type GetJourneysResponse struct {
	Journeys []Journey `json:"journeys"`
}
