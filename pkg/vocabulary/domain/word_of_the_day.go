package domain

import "github.com/namhq1989/vocab-booster-utilities/language"

type WordOfTheDay struct {
	Vocabulary  VocabularyBrief
	Information language.Multilingual
}
