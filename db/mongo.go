package db

import (
	"context"
	"portofolio/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbClient() (*mongo.Database, error) {
	clientOpt := options.Client()
	clientOpt.ApplyURI("mongodb://" + config.Configuration().Database.Mongodb.Url)
	client, err := mongo.NewClient(clientOpt)

	var ctx = context.Background()

	timeout := 3 * time.Second

	var ctxTo, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctxTo)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctxTo, nil)

	if err != nil {
		panic(err)
	}

	return client.Database(config.Configuration().Database.Mongodb.Collection), nil
}
