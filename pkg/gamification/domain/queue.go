package domain

type QueueExerciseAnsweredPoint struct {
	UserID         string
	ExerciseID     string
	Point          int64
	CompletionTime int
}

type QueueVocabularySentenceContributedPoint struct {
	UserID         string
	VocabularyID   string
	Point          int64
	CompletionTime int
}
