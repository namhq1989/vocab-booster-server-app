package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
)

type CommunitySentenceMapper struct{}

func (CommunitySentenceMapper) FromGrpcToDomain(sentence *vocabularypb.CommunitySentence) (*domain.CommunitySentence, error) {
	result := domain.CommunitySentence{
		ID:                   sentence.GetId(),
		VocabularyID:         sentence.GetVocabularyId(),
		Content:              dto.ConvertGrpcDataToMultilingual(sentence.GetContent()),
		RequiredVocabularies: sentence.GetRequiredVocabularies(),
		RequiredTense:        sentence.GetRequiredTense(),
		Clauses:              make([]domain.SentenceClause, 0),
		PosTags:              make([]domain.PosTag, 0),
		Sentiment: domain.Sentiment{
			Polarity:     sentence.GetSentiment().GetPolarity(),
			Subjectivity: sentence.GetSentiment().GetSubjectivity(),
		},
		Dependencies: make([]domain.Dependency, 0),
		Verbs:        make([]domain.Verb, 0),
		Level:        sentence.GetLevel(),
		StatsLike:    int(sentence.GetStatsLike()),
		IsLiked:      sentence.GetIsLiked(),
		CreatedAt:    sentence.GetCreatedAt().AsTime(),
	}

	for _, clause := range sentence.GetClauses() {
		result.Clauses = append(result.Clauses, domain.SentenceClause{
			Clause:         clause.GetClause(),
			Tense:          clause.GetTense(),
			IsPrimaryTense: clause.GetIsPrimaryTense(),
		})
	}

	for _, pos := range sentence.GetPosTags() {
		result.PosTags = append(result.PosTags, domain.PosTag{
			Word:  pos.GetWord(),
			Value: pos.GetValue(),
			Level: int(pos.GetLevel()),
		})
	}

	for _, dep := range sentence.GetDependencies() {
		result.Dependencies = append(result.Dependencies, domain.Dependency{
			Word:   dep.GetWord(),
			DepRel: dep.GetDepRel(),
			Head:   dep.GetHead(),
		})
	}

	for _, verb := range sentence.GetVerbs() {
		result.Verbs = append(result.Verbs, domain.Verb{
			Base:                verb.GetBase(),
			Past:                verb.GetPast(),
			PastParticiple:      verb.GetPastParticiple(),
			Gerund:              verb.GetGerund(),
			ThirdPersonSingular: verb.GetThirdPersonSingular(),
		})
	}

	return &result, nil
}

func (CommunitySentenceMapper) FromGrpcToDomainBrief(sentence *vocabularypb.CommunitySentenceBrief) (*domain.CommunitySentenceBrief, error) {
	result := domain.CommunitySentenceBrief{
		ID:        sentence.GetId(),
		Content:   dto.ConvertGrpcDataToMultilingual(sentence.GetContent()),
		Level:     sentence.GetLevel(),
		StatsLike: int(sentence.GetStatsLike()),
		IsLiked:   sentence.GetIsLiked(),
	}

	return &result, nil
}
