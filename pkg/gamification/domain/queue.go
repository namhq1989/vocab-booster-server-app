package domain

type QueueAddAnswerExercisePoint struct {
	UserID     string
	ExerciseID string
	Point      int64
}

type QueueAddContributeVocabularySentencePoint struct {
	UserID       string
	VocabularyID string
	Point        int64
}
