package domain

import "github.com/namhq1989/vocab-booster-utilities/language"

type PosTag struct {
	Word  string
	Value string
	Level int
}

type Sentiment struct {
	Polarity     float64
	Subjectivity float64
}

type Dependency struct {
	Word   string
	DepRel string
	Head   string
}

type Verb struct {
	Base                string
	Past                string
	PastParticiple      string
	Gerund              string
	ThirdPersonSingular string
}

type SentenceClause struct {
	Clause         string
	Tense          string
	IsPrimaryTense bool
}

type SentenceGrammarError struct {
	Message     language.Multilingual
	Segment     string
	Replacement string
}
