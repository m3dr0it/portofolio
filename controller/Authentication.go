package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"portofolio/model"
	"portofolio/service"
	"portofolio/util"
	"strings"

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
		})
		w.Write(responseJson)
		return
	}

	token, errGenerateToken := service.GenerateJWT(userLogin.Username)

	if errGenerateToken != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "failed",
		})

		w.Write(responseJson)
		return
	}

	responseJson, _ := json.Marshal(model.BaseResponse{
		Message: "success",
		Data: model.LoginResponse{
			AccessToken: token,
		},
	})
	w.Write(responseJson)
}

func WhoAmi(w http.ResponseWriter, r *http.Request) {
	allowedMethod := []string{"GET"}
	w.Header().Add("Content-Type", "application/json")

	if !slices.Contains(allowedMethod, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	auth := r.Header.Get("Authorization")

	if auth == "" {
		response, _ := json.Marshal(model.BaseResponse{
			Message: "Header Authorization Required",
		})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}
	auth = strings.Replace(auth, "Bearer ", "", -1)
	userInfo, err := service.ValidateJWT(auth)

	if err != nil {
		response, _ := json.Marshal(model.BaseResponse{
			Message: "Invalid Token or Expired",
		})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(model.BaseResponse{
		Message: "success",
		Data:    userInfo,
	})
	w.Write(response)

}
