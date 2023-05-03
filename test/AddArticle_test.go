package test

import (
	"log"
	"portofolio/model"
	"portofolio/service"
	"testing"
	"time"
)

func Test_AddArticle(t *testing.T) {
	article1 := model.Article{
		Id:        "nonono",
		Title:     "Test 1",
		CreatedAt: time.Now(),
		Content:   `Testots`,
	}

	err := service.AddArticle(article1)

	if err != nil {
		log.Println(err.Error())
	}

}
