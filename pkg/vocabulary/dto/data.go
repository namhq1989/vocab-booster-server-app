package dto

import "github.com/namhq1989/vocab-booster-utilities/language"

type PosTag struct {
	Word  string `json:"word"`
	Value string `json:"value"`
	Level int    `json:"level"`
}

type Sentiment struct {
	Polarity     float64 `json:"polarity"`
	Subjectivity float64 `json:"subjectivity"`
}

type Dependency struct {
	Word   string `json:"word"`
	DepRel string `json:"depRel"`
	Head   string `json:"head"`
}

type Verb struct {
	Base                string `json:"base"`
	Past                string `json:"past"`
	PastParticiple      string `json:"pastParticiple"`
	Gerund              string `json:"gerund"`
	ThirdPersonSingular string `json:"thirdPersonSingular"`
}

type SentenceClause struct {
	Clause         string `json:"clause"`
	Tense          string `json:"tense"`
	IsPrimaryTense bool   `json:"isPrimaryTense"`
}

type SentenceGrammarError struct {
	Message     language.Multilingual `json:"message"`
	Segment     string                `json:"segment"`
	Replacement string                `json:"replacement"`
}
