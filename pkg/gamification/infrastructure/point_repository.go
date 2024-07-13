package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
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
