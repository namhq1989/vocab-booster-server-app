package dto

type GetUserBookmarkedVocabulariesRequest struct {
	PageToken string `json:"pageToken"`
}

type GetUserBookmarkedVocabulariesResponse struct {
	Vocabularies  []VocabularyBrief `json:"vocabularies"`
	NextPageToken string            `json:"nextPageToken"`
}
