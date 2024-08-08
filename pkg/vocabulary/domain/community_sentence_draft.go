package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceDraft struct {
	ID                   string
	Content              language.Multilingual
	RequiredVocabularies []string
	RequiredTense        string
	IsCorrect            bool
	ErrorCode            string
	GrammarErrors        []SentenceGrammarError
	Sentiment            Sentiment
	Clauses              []SentenceClause
	Level                string
	CreatedAt            time.Time
}
