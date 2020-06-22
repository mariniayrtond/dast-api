package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
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
	return nil, nil
}
