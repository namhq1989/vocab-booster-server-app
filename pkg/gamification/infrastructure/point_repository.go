package infrastructure

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/infrastructure/aggregatemodel"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/infrastructure/dbmodel"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PointRepository struct {
	db             *database.Database
	collectionName string
}

func NewPointRepository(db *database.Database) PointRepository {
	r := PointRepository{
		db:             db,
		collectionName: database.Collections.GamificationPoint,
	}
	r.ensureIndexes()
	return r
}

func (r PointRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r PointRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r PointRepository) CreatePoint(ctx *appcontext.AppContext, point domain.Point) error {
	doc, err := dbmodel.Point{}.FromDomain(point)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r PointRepository) AggregateUserPointsInTimeRange(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]domain.UserAggregatedPoint, error) {
	var result = make([]domain.UserAggregatedPoint, 0)

	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return result, err
	}

	var pipelines = mongo.Pipeline{}

	pipelines = append(pipelines,
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "userId", Value: uid,
					},
					{
						Key: "createdAt", Value: bson.M{
							"$gte": from,
							"$lte": to,
						},
					},
				},
			},
		},
		bson.D{
			{
				Key: "$group", Value: bson.D{
					{
						Key: "_id", Value: bson.D{
							{
								Key: "$dateToString", Value: bson.M{
									"format":   "%d/%m",
									"date":     "$createdAt",
									"timezone": timezone,
								},
							},
						},
					},
					{
						Key: "point", Value: bson.M{
							"$sum": "$point",
						},
					},
					{
						Key: "createdAt", Value: bson.M{
							"$first": "$createdAt",
						},
					},
				},
			},
		},
		bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "_id", Value: 1},
					{Key: "point", Value: 1},
					{Key: "createdAt", Value: 1},
				},
			},
		},
		bson.D{
			{
				Key: "$sort", Value: bson.M{"createdAt": 1},
			},
		},
	)

	cursor, err := r.collection().Aggregate(ctx.Context(), pipelines)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []aggregatemodel.UserPoint
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, nil
	}

	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		dateStr := manipulation.FormatDDMM(d)
		docIndex := slices.IndexFunc(docs, func(point aggregatemodel.UserPoint) bool {
			return point.Date == dateStr
		})
		var point int64 = 0
		if docIndex != -1 {
			point = docs[docIndex].Point
		}

		result = append(result, domain.UserAggregatedPoint{
			Date:  dateStr,
			Point: point,
		})
	}

	return result, nil
}
