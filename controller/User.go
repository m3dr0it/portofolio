package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"portofolio/model"
	"portofolio/service"
	"portofolio/util"

	"golang.org/x/exp/slices"
)

func Login(w http.ResponseWriter, r *http.Request) {
	allowedMethod := []string{"POST"}
	message := ""
	w.Header().Add("Content-Type", "application/json")

	responseJson, err := json.Marshal(model.BaseResponse{
		Message: message,
		Data:    "",
	})

	var userLoginStuct model.UserLogin

	if !util.ValidatePayload(r, userLoginStuct) {
		w.WriteHeader(http.StatusBadRequest)
		responseJson, _ := json.Marshal(model.BaseResponse{
			Message: "missing required field",
			Data:    "",
		})
		w.Write(responseJson)
		return
	}

	if !slices.Contains(allowedMethod, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userLogin := new(model.UserLogin)
	errDecoder := json.NewDecoder(r.Body).Decode(userLogin)
	if errDecoder != nil {
		log.Println(err.Error())
	}

	signed := service.UserLogin(*userLogin)

	if !signed {
		message = "failed"
	} else {
		message = "success"
	}

	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write(responseJson)
}
