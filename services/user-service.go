package services

import (
	"errors"
	"fmt"
	apiModels "github.com/avnigenc/go-api/models/api"
	businessModels "github.com/avnigenc/go-api/models/business"
	userRepo "github.com/avnigenc/go-api/repository"
	"github.com/avnigenc/go-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authenticate(loginModel apiModels.LoginRequestModel) (string, error) {
	var err error
	user, err := userRepo.GetByEmail(loginModel.Email)

	if err != nil {
		fmt.Println("[UserService]: user not found: ", loginModel.Email)
		return "", errors.New("user not found")
	}

	if utils.CheckPasswordHash(loginModel.Password, user.Password) {
		fmt.Println("[UserService]: password wrong: ", loginModel.Email)
		return "", errors.New("password wrong")
	}

	token, _ := utils.GenerateAccessToken(user.Id)
	return token, nil
}

func CreateUser(registerModel apiModels.RegisterRequestModel) (businessModels.User, error) {
	var err error

	_, err = userRepo.GetByEmail(registerModel.Email)
	if err == nil {
		fmt.Println("[UserService]: user already exists: ", registerModel.Email)
		return businessModels.User{}, errors.New("user already exists")
	}

	hashedPassword, _ := utils.HashPassword(registerModel.Password)
	user := businessModels.User {
		Email: 		registerModel.Email,
		Password:	hashedPassword,
		FirstName:	registerModel.FirstName,
		LastName: 	registerModel.LastName,
	}

	createdUser, err := userRepo.Insert(user)
	if err != nil {
		fmt.Println("[UserService]: create user error")
		return businessModels.User{}, errors.New("create user error")
	}

	user.Id = createdUser.InsertedID.(primitive.ObjectID)

	return user, nil
}

func GetUserById(userId string) (businessModels.User, error) {
	var err error

	user, err := userRepo.GetById(userId)
	if err != nil {
		fmt.Println("[UserService]: user not found: ", userId)
		return businessModels.User{}, errors.New("user not found")
	}

	return user, nil
}
