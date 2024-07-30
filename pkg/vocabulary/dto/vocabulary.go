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
	Examples      []VocabularyExample    `json:"examples"`
	IsBookmarked  bool                   `json:"isBookmarked"`
}

type VocabularyDefinition struct {
	Pos        string                `json:"pos"`
	Definition language.Multilingual `json:"definition"`
}

type VocabularyMainWord struct {
	Word string `json:"word"`
	Base string `json:"base"`
	Pos  string `json:"pos"`
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
		IsBookmarked:  vocab.IsBookmarked,
	}

	for _, def := range vocab.Definitions {
		result.Definitions = append(result.Definitions, VocabularyDefinition{
			Pos:        def.Pos,
			Definition: def.Definition.GetLocalized(lang),
		})
	}

	for _, example := range vocab.Examples {
		result.Examples = append(result.Examples, VocabularyExample{
			ID:      example.ID,
			Audio:   example.Audio,
			Content: example.Content.GetLocalized(lang),
			MainWord: VocabularyMainWord{
				Word: example.MainWord.Word,
				Base: example.MainWord.Base,
				Pos:  example.MainWord.Pos,
			},
		})
	}

	return result
}

type VocabularyBrief struct {
	ID            string   `json:"id"`
	Term          string   `json:"term"`
	PartsOfSpeech []string `json:"partsOfSpeech"`
	Ipa           string   `json:"ipa"`
	Audio         string   `json:"audio"`
}

func (VocabularyBrief) FromDomain(vocab domain.VocabularyBrief) VocabularyBrief {
	result := VocabularyBrief{
		ID:            vocab.ID,
		Term:          vocab.Term,
		PartsOfSpeech: vocab.PartsOfSpeech,
		Ipa:           vocab.Ipa,
		Audio:         vocab.Audio,
	}

	return result
}
