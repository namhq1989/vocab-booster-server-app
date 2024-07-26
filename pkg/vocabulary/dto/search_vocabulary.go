package dto

type SearchVocabularyRequest struct {
	Term string `query:"term" validate:"required" message:"vocabulary_invalid_term"`
}

type SearchVocabularyResponse struct {
	Found       bool        `json:"found"`
	Vocabulary  *Vocabulary `json:"vocabulary"`
	Suggestions []string    `json:"suggestions"`
}
