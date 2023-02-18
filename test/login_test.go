package test

import (
	"encoding/json"
	"log"
	"net/http"
	"portofolio/model"
	"testing"
)

var baseUrl string = "http://localhost:8000/api/v1"

func TestLogin(t *testing.T) {
	url := baseUrl + "/login"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
	}

	baseResponse := new(model.BaseResponse)
	err1 := json.NewDecoder(res.Body).Decode(&baseResponse)

	if err1 != nil {
		log.Println(err1.Error())
	}

	if baseResponse.Message != "success" {
		log.Println("Login failed")
		t.Fail()
	}
}
