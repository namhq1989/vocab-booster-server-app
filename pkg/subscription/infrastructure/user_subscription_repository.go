package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
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
			{
				Keys: bson.D{{Key: "isPremium", Value: 1}, {Key: "endAt", Value: 1}},
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

func (r UserSubscriptionRepository) FindUserSubscriptionByUserID(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.UserSubscription
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"userId": uid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserSubscriptionRepository) UpsertUserSubscription(ctx *appcontext.AppContext, us domain.UserSubscription) error {
	doc, err := dbmodel.UserSubscription{}.FromDomain(us)
	if err != nil {
		return err
	}

	_, err = r.collection().ReplaceOne(ctx.Context(), bson.M{"userId": doc.UserID}, doc, options.Replace().SetUpsert(true))
	return err
}

func (r UserSubscriptionRepository) FindExpiredUserSubscriptionsByDate(ctx *appcontext.AppContext, date time.Time) ([]domain.UserSubscription, error) {
	var (
		condition = bson.M{
			"isPremium": true,
			"endAt":     bson.M{"$lte": date},
		}
		result = make([]domain.UserSubscription, 0)
	)

	cursor, err := r.collection().Find(ctx.Context(), condition, nil)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.UserSubscription
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
