package dto

type GetCommunitySentencesRequest struct {
	PageToken string `query:"pageToken"`
}

type GetCommunitySentencesResponse struct {
	Sentences     []CommunitySentenceBrief `json:"sentences"`
	NextPageToken string                   `json:"nextPageToken"`
}
