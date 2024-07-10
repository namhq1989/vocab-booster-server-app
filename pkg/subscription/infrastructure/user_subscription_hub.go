package infrastructure

import (
	"errors"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/infrastructure/dbmodel"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSubscriptionHub struct {
	db             *database.Database
	collectionName string
}

func NewUserSubscriptionHub(db *database.Database) UserSubscriptionHub {
	return UserSubscriptionHub{
		db:             db,
		collectionName: database.Collections.UserSubscription,
	}
}

func (r UserSubscriptionHub) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserSubscriptionHub) FindUserSubscriptionByUserID(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
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

func (r UserSubscriptionHub) CreateUserSubscription(ctx *appcontext.AppContext, us domain.UserSubscription) error {
	doc, err := dbmodel.UserSubscription{}.FromDomain(us)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
