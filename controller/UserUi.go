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

func LoginUI(w http.ResponseWriter, r *http.Request) {
	allowedMethod := []string{"POST"}

	if !slices.Contains(allowedMethod, r.Method) {
		return
	}

	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var reqBody map[string]string
	errParseReBody := json.Unmarshal(body, &reqBody)

	if errParseReBody != nil {
		return
	}

	var userLoginStuct model.UserLogin
	if !util.ValidatePayload(reqBody, userLoginStuct) {
		return
	}

	userLogin := model.UserLogin{
		Username: reqBody["username"],
		Password: reqBody["password"],
	}

	signed := service.UserLogin(userLogin)
	if signed {

	}
	return
}
