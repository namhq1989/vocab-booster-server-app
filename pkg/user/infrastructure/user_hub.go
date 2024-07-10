package infrastructure

import (
	"errors"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/infrastructure/dbmodel"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHub struct {
	db             *database.Database
	collectionName string
}

func NewUserHub(db *database.Database) UserHub {
	return UserHub{
		db:             db,
		collectionName: database.Collections.User,
	}
}

func (r UserHub) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserHub) FindUserByEmail(ctx *appcontext.AppContext, email string) (*domain.User, error) {
	var doc dbmodel.User
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"email": email,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserHub) FindUserByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.User
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": uid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserHub) CreateUser(ctx *appcontext.AppContext, user domain.User) error {
	doc, err := dbmodel.User{}.FromDomain(user)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
