package dto

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentence struct {
	ID                 string                `json:"id"`
	VocabularyID       string                `json:"vocabularyId"`
	Content            language.Multilingual `json:"content"`
	RequiredVocabulary []string              `json:"requiredVocabulary"`
	RequiredTense      string                `json:"requiredTense"`
	Clauses            []SentenceClause      `json:"clauses"`
	PosTags            []PosTag              `json:"posTags"`
	Sentiment          Sentiment             `json:"sentiment"`
	Dependencies       []Dependency          `json:"dependencies"`
	Verbs              []Verb                `json:"verbs"`
	Level              string                `json:"level"`
	StatsLike          int                   `json:"statsLike"`
	IsLiked            bool                  `json:"isLiked"`
}

func (CommunitySentence) FromDomain(sentence domain.CommunitySentence, lang string) CommunitySentence {
	var clauses = make([]SentenceClause, 0)
	for _, clause := range sentence.Clauses {
		clauses = append(clauses, SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense,
			IsPrimaryTense: clause.IsPrimaryTense,
		})
	}

	var posTags = make([]PosTag, 0)
	for _, pos := range sentence.PosTags {
		posTags = append(posTags, PosTag{
			Word:  pos.Word,
			Value: pos.Value,
			Level: pos.Level,
		})
	}

	var deps = make([]Dependency, 0)
	for _, dep := range sentence.Dependencies {
		deps = append(deps, Dependency{
			Word:   dep.Word,
			DepRel: dep.DepRel,
			Head:   dep.Head,
		})
	}

	var verbs = make([]Verb, 0)
	for _, verb := range sentence.Verbs {
		verbs = append(verbs, Verb{
			Base:                verb.Base,
			Past:                verb.Past,
			PastParticiple:      verb.PastParticiple,
			Gerund:              verb.Gerund,
			ThirdPersonSingular: verb.ThirdPersonSingular,
		})
	}

	return CommunitySentence{
		ID:                 sentence.ID,
		VocabularyID:       sentence.VocabularyID,
		Content:            sentence.Content.GetLocalized(lang),
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      sentence.RequiredTense,
		Clauses:            clauses,
		PosTags:            posTags,
		Sentiment: Sentiment{
			Polarity:     sentence.Sentiment.Polarity,
			Subjectivity: sentence.Sentiment.Subjectivity,
		},
		Dependencies: deps,
		Verbs:        verbs,
		Level:        sentence.Level,
		StatsLike:    sentence.StatsLike,
		IsLiked:      sentence.IsLiked,
	}
}

type CommunitySentenceBrief struct {
	ID        string                `json:"id"`
	Content   language.Multilingual `json:"content"`
	Level     string                `json:"level"`
	StatsLike int                   `json:"statsLike"`
	IsLiked   bool                  `json:"isLiked"`
}

func (CommunitySentenceBrief) FromDomain(sentence domain.CommunitySentenceBrief, lang string) CommunitySentenceBrief {
	return CommunitySentenceBrief{
		ID:        sentence.ID,
		Content:   sentence.Content.GetLocalized(lang),
		Level:     sentence.Level,
		StatsLike: sentence.StatsLike,
		IsLiked:   sentence.IsLiked,
	}
}
