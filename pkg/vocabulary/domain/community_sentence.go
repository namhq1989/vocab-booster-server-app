package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentence struct {
	ID                   string
	VocabularyID         string
	Content              language.Multilingual
	MainWord             VocabularyMainWord
	RequiredVocabularies []string
	RequiredTense        string
	Clauses              []SentenceClause
	PosTags              []PosTag
	Sentiment            Sentiment
	Dependencies         []Dependency
	Verbs                []Verb
	Level                string
	StatsLike            int
	IsLiked              bool
	CreatedAt            time.Time
}

type CommunitySentenceBrief struct {
	ID        string
	Content   language.Multilingual
	Level     string
	StatsLike int
	IsLiked   bool
}
