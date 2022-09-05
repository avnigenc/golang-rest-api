package repository

import (
	"errors"

	business "github.com/avnigenc/go-api/models/business"
	modules "github.com/avnigenc/go-api/modules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(user business.User) (*mongo.InsertOneResult, error) {
	collection, ctx := modules.UserCollection()

	userResult, err := collection.InsertOne(ctx, bson.D{
		{Key: "Email", Value: user.Email},
		{Key: "Password", Value: user.Password},
		{Key: "FirstName", Value: user.FirstName},
		{Key: "LastName", Value: user.LastName},
	})
	if err == nil {
		return userResult, nil
	}
	return nil, nil
}

func GetById(userId string) (business.User, error) {
	collection, ctx := modules.UserCollection()

	var user business.User
	objID, _ := primitive.ObjectIDFromHex(userId)
	err := collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return business.User{}, errors.New("user not found")
	}
	return user, err
}

func GetByEmail(email string) (business.User, error) {
	collection, ctx := modules.UserCollection()

	var user business.User
	err := collection.FindOne(ctx, bson.D{
		{Key: "Email", Value: email},
	}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return business.User{}, errors.New("user not found")
	}
	return user, err
}
