package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type Vocabulary struct {
	ID            string                 `json:"id"`
	AuthorID      string                 `json:"authorId"`
	Term          string                 `json:"term"`
	Definitions   []VocabularyDefinition `json:"definitions"`
	PartsOfSpeech []string               `json:"partsOfSpeech"`
	Ipa           string                 `json:"ipa"`
	Audio         string                 `json:"audio"`
	Synonyms      []string               `json:"synonyms"`
	Antonyms      []string               `json:"antonyms"`
}

type VocabularyDefinition struct {
	Pos        string                `json:"pos"`
	Definition language.Multilingual `json:"definition"`
}

func (Vocabulary) FromDomain(vocab domain.Vocabulary, lang string) Vocabulary {
	result := Vocabulary{
		ID:            vocab.ID,
		AuthorID:      vocab.AuthorID,
		Term:          vocab.Term,
		Definitions:   make([]VocabularyDefinition, 0),
		PartsOfSpeech: vocab.PartsOfSpeech,
		Ipa:           vocab.Ipa,
		Audio:         vocab.Audio,
		Synonyms:      vocab.Synonyms,
		Antonyms:      vocab.Antonyms,
	}

	for _, def := range vocab.Definitions {
		result.Definitions = append(result.Definitions, VocabularyDefinition{
			Pos:        def.Pos,
			Definition: def.Definition.GetLocalized(lang),
		})
	}

	return result
}
