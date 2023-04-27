package db

import (
	"context"
	"portofolio/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbClient() (*mongo.Database, error) {
	clientOpt := options.Client()
	clientOpt.ApplyURI("mongodb://" + config.Configuration().Database.Mongodb.Url)
	client, err := mongo.NewClient(clientOpt)

	var ctx = context.Background()

	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client.Database(config.Configuration().Database.Mongodb.Collection), nil
}
