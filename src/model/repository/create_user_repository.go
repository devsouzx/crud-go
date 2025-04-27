package repository

import (
	"context"
	"os"

	"github.com/devsouzx/crud-go/src/configuration/logger"
	"github.com/devsouzx/crud-go/src/configuration/rest_err"
	"github.com/devsouzx/crud-go/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	collectionName := os.Getenv("MONGODB_USER_DB")

	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomain.GetJsonValue()
	if err != nil {
		logger.Error("Error trying to convert user domain to json",
			err,
			zap.String("journey", "createUser"))	
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}