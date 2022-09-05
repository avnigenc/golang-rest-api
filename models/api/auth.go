package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginRequestModel struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type LoginResponseModel struct {
	Token string
}

type RegisterRequestModel struct {
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	Email     string `json:"Email" binding:"required"`
	Password  string `json:"Password" binding:"required"`
}

type RegisterResponseModel struct {
	Id        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
}
