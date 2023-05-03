package model

import "time"

type Article struct {
	Id        string    `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Content   string    `json:"content" bson:"content"`
}
