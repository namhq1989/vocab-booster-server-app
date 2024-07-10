package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User        string
	UserJourney string

	UserSubscription        string
	UserSubscriptionHistory string
}{
	User:        "user.users",
	UserJourney: "users.journeys",

	UserSubscription:        "subscription.userSubscriptions",
	UserSubscriptionHistory: "subscription.userSubscriptionHistories",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
