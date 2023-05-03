package service

import (
	"context"
	"log"
	"portofolio/config"
	"portofolio/db"
	"portofolio/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetArticles() ([]model.Article, error) {
	ctx := context.Background()
	db, err := db.MongoDbClient()

	if err != nil {
		log.Println(err.Error())
	}

	articles, err := db.Collection(config.Configuration().Database.Mongodb.Collection).Find(ctx, bson.M{})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var resultArticles []model.Article

	for articles.Next(ctx) {
		var article model.Article

		if err := articles.Decode(&article); err != nil {
			panic(err)
		}

		resultArticles = append(resultArticles, article)
	}

	return resultArticles, nil
}

func AddArticle(article model.Article) error {
	ctx := context.Background()
	db, err := db.MongoDbClient()

	if err != nil {
		log.Println(err.Error())
	}

	result, err := db.Collection(config.Configuration().Database.Mongodb.Collection).InsertOne(ctx, article)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println(result)

	return nil
}
