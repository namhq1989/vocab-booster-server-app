package queue

var TypeNames = struct {
	ScanExpiredUserSubscription string
	DowngradeUserSubscription   string

	ExerciseAnswered string

	GamificationAddAnswerExercisePoint               string
	GamificationAddContributeVocabularySentencePoint string
}{
	ScanExpiredUserSubscription: "subscription.scanExpiredUserSubscription",
	DowngradeUserSubscription:   "subscription.downgradeUserSubscription",

	ExerciseAnswered: "exercise.exerciseAnswered",

	GamificationAddAnswerExercisePoint:               "gamification.addAnswerExercisePoint",
	GamificationAddContributeVocabularySentencePoint: "gamification.addContributeVocabularySentencePoint",
}
