package services

import (
	"context"
	"errors"

	"example.com/internal/models"
	"example.com/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(ctx context.Context, email, password string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{Email: email, Password: hashed}
	_, err = utils.DB.Collection("users").InsertOne(ctx, user)
	return err
}

func Login(ctx context.Context, email, password string) (string, error) {
	var user models.User

	err := utils.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("not found in db")
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}
