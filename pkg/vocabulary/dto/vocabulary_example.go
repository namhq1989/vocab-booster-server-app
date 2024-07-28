package dto

import "github.com/namhq1989/vocab-booster-utilities/language"

type VocabularyExample struct {
	ID       string                `json:"id"`
	Audio    string                `json:"audio"`
	Content  language.Multilingual `json:"content"`
	MainWord VocabularyMainWord    `json:"mainWord"`
}
