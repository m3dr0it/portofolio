package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"portofolio/controller"
	"portofolio/service"
	"portofolio/test/model"
	"testing"
)

func Test_TestLoginEndpoint(t *testing.T) {
	reqBody := []byte(`{"username": "John", "password": "root"}`)
	req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(reqBody))
	response := httptest.NewRecorder()
	controller.Login(response, req)

	var respBody model.LoginResponseTest
	errParseReBody := json.Unmarshal(response.Body.Bytes(), &respBody)

	if errParseReBody != nil {
		t.Fail()
	}

	data := respBody.Data
	log.Println(data.AccessToken)

	_, errCekToken := service.ValidateJWT(data.AccessToken)

	if errCekToken != nil {
		t.Fail()
	}

	if response.Code != 200 {
		t.Fail()
	}
}
