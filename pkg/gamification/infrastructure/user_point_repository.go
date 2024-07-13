package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/infrastructure/dbmodel"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserPointRepository struct {
	db             *database.Database
	collectionName string
}

func NewUserPointRepository(db *database.Database) UserPointRepository {
	r := UserPointRepository{
		db:             db,
		collectionName: database.Collections.GamificationUserPoint,
	}
	r.ensureIndexes()
	return r
}

func (r UserPointRepository) ensureIndexes() {
	var (
		ctx            = context.Background()
		opts           = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isUniqueUserID = true
		indexes        = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isUniqueUserID,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r UserPointRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserPointRepository) FindUserPoint(ctx *appcontext.AppContext, userID string) (*domain.UserPoint, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.UserPoint
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"userId": uid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}

func (r UserPointRepository) IncreasePoint(ctx *appcontext.AppContext, userID string, point int64) error {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return apperrors.User.InvalidUserID
	}

	_, err = r.collection().UpdateOne(ctx.Context(), bson.M{"userId": uid}, bson.M{"$inc": bson.M{"point": point}}, options.Update().SetUpsert(true))
	return err
}
