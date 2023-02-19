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

	message := "failed"
	if signed {
		message = "success"
	}

	w.Header().Add("Content-Type", "application/json")
	responseJson, _ := json.Marshal(model.BaseResponse{
		Message: message,
		Data:    "",
	})
	w.Write(responseJson)
}
