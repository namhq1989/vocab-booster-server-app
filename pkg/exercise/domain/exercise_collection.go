package domain

type ExerciseCollection struct {
	ID              string
	Name            string
	Slug            string
	Translated      string
	StatsExercises  int
	StatsInteracted int
	Image           string
}
