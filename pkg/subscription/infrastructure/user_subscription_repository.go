package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSubscriptionRepository struct {
	db             *database.Database
	collectionName string
}

func NewUserSubscriptionRepository(db *database.Database) UserSubscriptionRepository {
	r := UserSubscriptionRepository{
		db:             db,
		collectionName: database.Collections.UserSubscription,
	}
	r.ensureIndexes()
	return r
}

func (r UserSubscriptionRepository) ensureIndexes() {
	var (
		ctx            = context.Background()
		opts           = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isUserIDUnique = true
		indexes        = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isUserIDUnique,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r UserSubscriptionRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserSubscriptionRepository) UpsertUserSubscription(ctx *appcontext.AppContext, us domain.UserSubscription) error {
	doc, err := dbmodel.UserSubscription{}.FromDomain(us)
	if err != nil {
		return err
	}

	_, err = r.collection().ReplaceOne(ctx.Context(), bson.M{"userId": doc.UserID}, doc, options.Replace().SetUpsert(true))
	return err
}