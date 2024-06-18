package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string
}{
	User: "user.users",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
