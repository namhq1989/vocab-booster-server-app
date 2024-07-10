package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/infrastructure/dbmodel"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JourneyRepository struct {
	db             *database.Database
	collectionName string
}

func NewJourneyRepository(db *database.Database) JourneyRepository {
	r := JourneyRepository{
		db:             db,
		collectionName: database.Collections.UserJourney,
	}
	r.ensureIndexes()
	return r
}

func (r JourneyRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r JourneyRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r JourneyRepository) CreateJourney(ctx *appcontext.AppContext, journey domain.Journey) error {
	doc, err := dbmodel.Journey{}.FromDomain(journey)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r JourneyRepository) UpdateJourney(ctx *appcontext.AppContext, journey domain.Journey) error {
	doc, err := dbmodel.Journey{}.FromDomain(journey)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r JourneyRepository) FindUserCurrentJourney(ctx *appcontext.AppContext, userID string) (*domain.Journey, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.Journey
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"userId":     uid,
		"isLearning": true,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r JourneyRepository) FindJourneysByUserID(ctx *appcontext.AppContext, userID string) ([]domain.Journey, error) {
	result := make([]domain.Journey, 0)

	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return result, apperrors.User.InvalidUserID
	}

	cursor, err := r.collection().Find(ctx.Context(), bson.M{"userId": uid}, nil)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.Journey
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
