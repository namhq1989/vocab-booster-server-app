package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyHub interface {
	SearchVocabulary(ctx *appcontext.AppContext, performerID, term string) (*Vocabulary, []string, error)
	BookmarkVocabulary(ctx *appcontext.AppContext, userID, vocabularyID string) (bool, error)
	GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, userID, pageToken string) ([]VocabularyBrief, string, error)
	GetWordOfTheDay(ctx *appcontext.AppContext, lang string) (*WordOfTheDay, error)
	GetCommunitySentences(ctx *appcontext.AppContext, userID, vocabularyID, lang, pageToken string) ([]CommunitySentenceBrief, string, error)
	GetCommunitySentence(ctx *appcontext.AppContext, userID, sentenceID string) (*CommunitySentence, error)
	GetUserCommunitySentencesDraft(ctx *appcontext.AppContext, userID, vocabularyID, pageToken string) ([]CommunitySentenceDraft, string, error)
}

type VocabularyBrief struct {
	ID            string
	Term          string
	PartsOfSpeech []string
	Ipa           string
	Audio         string
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
	Examples      []VocabularyExample
	IsBookmarked  bool
}

type VocabularyDefinition struct {
	Pos        string
	Definition language.Multilingual
}

type VocabularyMainWord struct {
	Word string
	Base string
	Pos  string
}
