package main

import (
	"github.com/devsouzx/crud-go/src/controller"
	"github.com/devsouzx/crud-go/src/model/repository"
	"github.com/devsouzx/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}