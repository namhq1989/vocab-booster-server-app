package dto

type GetUserCommunitySentenceDraftsRequest struct {
	VocabularyID string `query:"vocabularyId"`
	PageToken    string `query:"pageToken"`
}

type GetUserCommunitySentenceDraftsResponse struct {
	Sentences     []CommunitySentenceDraft `json:"sentences"`
	NextPageToken string                   `json:"nextPageToken"`
}
