package dto

type BookmarkVocabularyRequest struct{}

type BookmarkVocabularyResponse struct {
	IsBookmarked bool `json:"isBookmarked"`
}
