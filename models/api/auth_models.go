package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginRequestModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseModel struct {
	Token string
}

type RegisterRequestModel struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type RegisterResponseModel struct {
	Id        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
}
