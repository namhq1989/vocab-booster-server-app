package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string

	UserSubscription string
}{
	User: "user.users",

	UserSubscription: "subscription.userSubscriptions",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
