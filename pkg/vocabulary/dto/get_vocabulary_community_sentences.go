package dto

type GetVocabularyCommunitySentencesRequest struct {
	PageToken string `query:"pageToken"`
}

type GetVocabularyCommunitySentencesResponse struct {
	Sentences     []Sentence `json:"sentences"`
	NextPageToken string     `json:"nextPageToken"`
}
