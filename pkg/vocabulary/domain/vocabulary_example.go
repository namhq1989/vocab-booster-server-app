package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyExample struct {
	ID       string
	Audio    string
	Content  language.Multilingual
	MainWord VocabularyMainWord
}
