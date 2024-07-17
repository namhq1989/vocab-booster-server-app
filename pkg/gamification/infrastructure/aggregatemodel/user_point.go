package aggregatemodel

import "time"

type UserPoint struct {
	Date      string    `bson:"_id"`
	Point     int64     `bson:"point"`
	CreatedAt time.Time `bson:"createdAt"`
}
