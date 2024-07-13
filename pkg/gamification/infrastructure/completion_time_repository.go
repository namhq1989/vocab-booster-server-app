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

type CompletionTimeRepository struct {
	db             *database.Database
	collectionName string
}

func NewCompletionTimeRepository(db *database.Database) CompletionTimeRepository {
	r := CompletionTimeRepository{
		db:             db,
		collectionName: database.Collections.GamificationCompletionTime,
	}
	r.ensureIndexes()
	return r
}

func (r CompletionTimeRepository) ensureIndexes() {
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

func (r CompletionTimeRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r CompletionTimeRepository) CreateCompletionTime(ctx *appcontext.AppContext, completionTime domain.CompletionTime) error {
	doc, err := dbmodel.CompletionTime{}.FromDomain(completionTime)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
