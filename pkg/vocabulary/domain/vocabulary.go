package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyHub interface {
	SearchVocabulary(ctx *appcontext.AppContext, performerID, term string) (*Vocabulary, []string, error)
}

type Vocabulary struct {
	ID            string
	AuthorID      string
	Term          string
	Definitions   []VocabularyDefinition
	PartsOfSpeech []string
	Ipa           string
	Audio         string
	Synonyms      []string
	Antonyms      []string
}

type VocabularyDefinition struct {
	Pos        string
	Definition language.Multilingual
}
