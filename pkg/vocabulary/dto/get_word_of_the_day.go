package dto

import "github.com/namhq1989/vocab-booster-utilities/language"

type GetWordOfTheDayRequest struct{}

type GetWordOfTheDayResponse struct {
	Vocabulary  VocabularyBrief       `json:"vocabulary"`
	Information language.Multilingual `json:"information"`
}
