package test

import (
	"log"
	"portofolio/service"
	"testing"
)

func Test_GetArticles(t *testing.T) {
	result, err := service.GetArticles()

	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}

	for _, res := range result {
		log.Println(res)
	}

}
