package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// docker run --name mongodb -p 27017:27017 -d mongo 
// go get go.mongodb.org/mongo-driver/mongo

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodbURI := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbDatabase), nil
}