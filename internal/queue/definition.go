package queue

var TypeNames = struct {
	ScanExpiredUserSubscription string
	DowngradeUserSubscription   string

	ExerciseAnswered string

	GamificationExerciseAnswered              string
	GamificationVocabularySentenceContributed string
}{
	ScanExpiredUserSubscription: "subscription.scanExpiredUserSubscription",
	DowngradeUserSubscription:   "subscription.downgradeUserSubscription",

	ExerciseAnswered: "exercise.exerciseAnswered",

	GamificationExerciseAnswered:              "gamification.exerciseAnswered",
	GamificationVocabularySentenceContributed: "gamification.vocabularySentenceContributed",
}
