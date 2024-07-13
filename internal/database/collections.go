package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string

	UserSubscription        string
	UserSubscriptionHistory string

	GamificationPoint          string
	GamificationCompletionTime string
	GamificationUserStats      string
}{
	User: "user.users",

	UserSubscription:        "subscription.userSubscriptions",
	UserSubscriptionHistory: "subscription.userSubscriptionHistories",

	GamificationPoint:          "gamification.points",
	GamificationCompletionTime: "gamification.completionTimes",
	GamificationUserStats:      "gamification.userStats",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
