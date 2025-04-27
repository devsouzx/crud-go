package main

import (
	"context"
	"log"

	"github.com/devsouzx/crud-go/src/configuration/database/mongodb"
	"github.com/devsouzx/crud-go/src/configuration/logger"
	"github.com/devsouzx/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect to MongoDB", err)	
		panic(err)
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}