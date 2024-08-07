package dto

type GetCommunitySentenceRequest struct{}

type GetCommunitySentenceResponse struct {
	Sentence CommunitySentence `json:"sentence"`
}
