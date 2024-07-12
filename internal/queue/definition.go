package queue

var TypeNames = struct {
	ScanExpiredUserSubscription string
	DowngradeUserSubscription   string

	ExerciseAnswered string
}{
	ScanExpiredUserSubscription: "subscription.scanExpiredUserSubscription",
	DowngradeUserSubscription:   "subscription.downgradeUserSubscription",

	ExerciseAnswered: "exercise.exerciseAnswered",
}
