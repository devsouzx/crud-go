package main

import (
	"context"
	"log"

	"github.com/devsouzx/crud-go/src/configuration/database/mongodb"
	"github.com/devsouzx/crud-go/src/configuration/logger"
	"github.com/devsouzx/crud-go/src/controller"
	"github.com/devsouzx/crud-go/src/controller/routes"
	"github.com/devsouzx/crud-go/src/model/repository"
	"github.com/devsouzx/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}