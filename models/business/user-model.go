package business_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `bson:"_id" json:"Id,omitempty"`
	FirstName	string
	LastName 	string
	Email 		string
	Password 	string
}