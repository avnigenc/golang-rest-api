package services

import (
	"errors"
	"fmt"

	models "github.com/avnigenc/go-api/models/api"
	business "github.com/avnigenc/go-api/models/business"
	repository "github.com/avnigenc/go-api/repository"
	"github.com/avnigenc/go-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(registerModel models.RegisterRequestModel) (business.User, error) {
	var err error

	_, err = repository.GetByEmail(registerModel.Email)
	if err == nil {
		fmt.Println("[UserService]: user already exists: ", registerModel.Email)
		return business.User{}, errors.New("user already exists")
	}

	hashedPassword, _ := utils.HashPassword(registerModel.Password)
	user := business.User{
		Email:     registerModel.Email,
		Password:  hashedPassword,
		FirstName: registerModel.FirstName,
		LastName:  registerModel.LastName,
	}

	createdUser, err := repository.Insert(user)
	if err != nil {
		fmt.Println("[UserService]: create user error")
		return business.User{}, errors.New("create user error")
	}

	user.Id = createdUser.InsertedID.(primitive.ObjectID)

	return user, nil
}

func GetUserById(userId string) (business.User, error) {
	var err error

	user, err := repository.GetById(userId)
	if err != nil {
		fmt.Println("[UserService]: user not found: ", userId)
		return business.User{}, errors.New("user not found")
	}

	return user, nil
}
