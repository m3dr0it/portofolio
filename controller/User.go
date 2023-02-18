package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"portofolio/model"

	"golang.org/x/exp/slices"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	allowedMethod := []string{"POST", "OPTION"}

	if !slices.Contains(allowedMethod, r.Method) {
		w.Write([]byte("Method Not Allowed"))
		return
	}

	userLogin := new(UserLogin)
	err := json.NewDecoder(r.Body).Decode(userLogin)
	if err != nil {
		log.Println(err.Error())
	}

	responseJson, err := json.Marshal(model.BaseResponse{
		Message: "success",
		Data:    "test",
	})

	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write(responseJson)
}
