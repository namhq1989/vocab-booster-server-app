package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceDraft struct {
	ID                   string                    `json:"id"`
	Content              language.Multilingual     `json:"content"`
	RequiredVocabularies []string                  `json:"requiredVocabularies"`
	RequiredTense        string                    `json:"requiredTense"`
	IsCorrect            bool                      `json:"isCorrect"`
	ErrorCode            string                    `json:"errorCode"`
	GrammarErrors        []SentenceGrammarError    `json:"grammarErrors"`
	Sentiment            Sentiment                 `json:"sentiment"`
	Clauses              []SentenceClause          `json:"clauses"`
	Level                string                    `json:"level"`
	CreatedAt            *httprespond.TimeResponse `json:"createdAt"`
}

func (CommunitySentenceDraft) FromDomain(sentence domain.CommunitySentenceDraft, lang string) CommunitySentenceDraft {
	var clauses = make([]SentenceClause, 0)
	for _, clause := range sentence.Clauses {
		clauses = append(clauses, SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense,
			IsPrimaryTense: clause.IsPrimaryTense,
		})
	}

	var grammarErrors = make([]SentenceGrammarError, 0)
	for _, err := range sentence.GrammarErrors {
		grammarErrors = append(grammarErrors, SentenceGrammarError{
			Message:     err.Message.GetLocalized(lang),
			Segment:     err.Segment,
			Replacement: err.Replacement,
		})
	}

	return CommunitySentenceDraft{
		ID:                   sentence.ID,
		Content:              sentence.Content.GetLocalized(lang),
		RequiredVocabularies: sentence.RequiredVocabularies,
		RequiredTense:        sentence.RequiredTense,
		Clauses:              clauses,
		IsCorrect:            sentence.IsCorrect,
		ErrorCode:            sentence.ErrorCode,
		Sentiment: Sentiment{
			Polarity:     sentence.Sentiment.Polarity,
			Subjectivity: sentence.Sentiment.Subjectivity,
		},
		GrammarErrors: grammarErrors,
		Level:         sentence.Level,
		CreatedAt:     httprespond.NewTimeResponse(sentence.CreatedAt),
	}
}
