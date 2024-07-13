package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User        string
	UserJourney string

	UserSubscription        string
	UserSubscriptionHistory string

	GamificationPoint     string
	GamificationUserPoint string
}{
	User:        "user.users",
	UserJourney: "users.journeys",

	UserSubscription:        "subscription.userSubscriptions",
	UserSubscriptionHistory: "subscription.userSubscriptionHistories",

	GamificationPoint:     "gamification.points",
	GamificationUserPoint: "gamification.userPoints",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
