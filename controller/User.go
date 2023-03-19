package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"portofolio/model"
	"portofolio/service"
	"portofolio/util"

	"golang.org/x/exp/slices"
)

func Login(w http.ResponseWriter, r *http.Request) {
	allowedMethod := []string{"POST"}

	if !slices.Contains(allowedMethod, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var reqBody map[string]string
	errParseReBody := json.Unmarshal(body, &reqBody)

	if errParseReBody != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "invalid payload",
		})
		w.Write(responseJson)
		return
	}

	var userLoginStuct model.UserLogin
	if !util.ValidatePayload(reqBody, userLoginStuct) {
		w.WriteHeader(http.StatusBadRequest)
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "missing required field",
		})
		w.Write(responseJson)
		return
	}

	userLogin := model.UserLogin{
		Username: reqBody["username"],
		Password: reqBody["password"],
	}

	signed := service.UserLogin(userLogin)
	w.Header().Add("Content-Type", "application/json")

	if !signed {
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "failed",
			Data:    model.Data{},
		})
		w.Write(responseJson)
		return
	}

	token, errGenerateToken := service.GenerateJWT(userLogin.Username)

	if errGenerateToken != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "failed",
			Data:    model.Data{},
		})

		w.Write(responseJson)
		return
	}

	responseJson, _ := json.Marshal(model.BaseResponse{
		Message: "success",
		Data: model.Data{
			AccessToken: token,
		},
	})
	w.Write(responseJson)

}
