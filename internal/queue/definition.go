package queue

var TypeNames = struct {
	ScanExpiredUserSubscription string
	DowngradeUserSubscription   string
}{
	ScanExpiredUserSubscription: "subscription.scanExpiredUserSubscription",
	DowngradeUserSubscription:   "subscription.downgradeUserSubscription",
}
