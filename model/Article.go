package model

import "time"

type Article struct {
	Id        string    `bson:"id"`
	Title     string    `bson:"title"`
	CreatedAt time.Time `bson:"createdAt"`
	Content   string    `bson:"content"`
}
