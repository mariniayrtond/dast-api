package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

type mongoUser struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func parseToMongoUser(user *model.User) mongoUser {
	return mongoUser{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func parseMongoMapToNativeUser(mh mongoUser) *model.User {
	return &model.User{
		ID:       mh.ID.Hex(),
		Name:     mh.Name,
		Email:    mh.Email,
		Password: mh.Password,
	}
}

func (u UserRepository) Save(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUser := parseToMongoUser(user)
	if _, err := u.collection.InsertOne(ctx, mongoUser); err != nil {
		return err
	}

	user.ID = mongoUser.ID.Hex()
	return nil
}

func (u UserRepository) Get(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filterCursor, err := u.collection.Find(ctx, bson.M{"_id": oID})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoUser
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeUser(episodesFiltered[0]), nil
}

func (u UserRepository) SearchByName(name string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := u.collection.Find(ctx, bson.M{"name": name})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoUser
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeUser(episodesFiltered[0]), nil
}
