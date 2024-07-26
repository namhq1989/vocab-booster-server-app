package domain

import "github.com/namhq1989/vocab-booster-utilities/language"

type ExerciseCollection struct {
	ID              string
	Name            language.Multilingual
	Slug            string
	StatsExercises  int
	StatsInteracted int
	Image           string
}
