package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoLogIn struct {
	UserID    string             `bson:"user_id"`
	Token     string             `bson:"token"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

func parseToMongoLogIn(l *model.LogIn) mongoLogIn {
	return mongoLogIn{
		UserID:    l.UserID,
		Token:     l.Token,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func parseMongoMapToNativeLogIn(mc mongoLogIn) *model.LogIn {
	return &model.LogIn{
		UserID: mc.UserID,
		Token:  mc.Token,
	}
}

type TokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(collection *mongo.Collection) *TokenRepository {
	return &TokenRepository{collection: collection}
}

func (t TokenRepository) Create(token *model.LogIn, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoLogIn := parseToMongoLogIn(token)
	if _, err := t.collection.InsertOne(ctx, mongoLogIn); err != nil {
		return err
	}

	return nil
}

func (t TokenRepository) Get(token string) (*model.LogIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := t.collection.Find(ctx, bson.M{"token": token})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoLogIn
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	if episodesFiltered == nil {
		return nil, nil
	}

	return parseMongoMapToNativeLogIn(episodesFiltered[0]), nil
}
