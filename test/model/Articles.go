package model

import "time"

type Article struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}
