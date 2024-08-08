package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
)

type CommunitySentenceDraftMapper struct{}

func (CommunitySentenceDraftMapper) FromGrpcToDomain(sentence *vocabularypb.CommunitySentenceDraft) (*domain.CommunitySentenceDraft, error) {
	result := domain.CommunitySentenceDraft{
		ID:                   sentence.GetId(),
		Content:              dto.ConvertGrpcDataToMultilingual(sentence.GetContent()),
		RequiredVocabularies: sentence.GetRequiredVocabularies(),
		RequiredTense:        sentence.GetRequiredTense(),
		IsCorrect:            sentence.IsCorrect,
		ErrorCode:            sentence.ErrorCode,
		GrammarErrors:        make([]domain.SentenceGrammarError, 0),
		Sentiment: domain.Sentiment{
			Polarity:     sentence.GetSentiment().GetPolarity(),
			Subjectivity: sentence.GetSentiment().GetSubjectivity(),
		},
		Clauses:   make([]domain.SentenceClause, 0),
		Level:     sentence.GetLevel(),
		CreatedAt: sentence.GetCreatedAt().AsTime(),
	}

	for _, clause := range sentence.GetClauses() {
		result.Clauses = append(result.Clauses, domain.SentenceClause{
			Clause:         clause.GetClause(),
			Tense:          clause.GetTense(),
			IsPrimaryTense: clause.GetIsPrimaryTense(),
		})
	}

	for _, err := range sentence.GetErrors() {
		result.GrammarErrors = append(result.GrammarErrors, domain.SentenceGrammarError{
			Message:     dto.ConvertGrpcDataToMultilingual(err.GetMessage()),
			Segment:     err.GetSegment(),
			Replacement: err.GetReplacement(),
		})
	}

	return &result, nil
}
