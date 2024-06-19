package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/infrastructure/dbmodel"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSubscriptionHistoryRepository struct {
	db             *database.Database
	collectionName string
}

func NewUserSubscriptionHistoryRepository(db *database.Database) UserSubscriptionHistoryRepository {
	r := UserSubscriptionHistoryRepository{
		db:             db,
		collectionName: database.Collections.UserSubscriptionHistory,
	}
	r.ensureIndexes()
	return r
}

func (r UserSubscriptionHistoryRepository) ensureIndexes() {
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

func (r UserSubscriptionHistoryRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserSubscriptionHistoryRepository) CreateUserSubscriptionHistory(ctx *appcontext.AppContext, history domain.UserSubscriptionHistory) error {
	doc, err := dbmodel.UserSubscriptionHistory{}.FromDomain(history)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
