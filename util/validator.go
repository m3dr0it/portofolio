package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func ValidatePayload(r *http.Request, model any) bool {
	numField := reflect.TypeOf(model).NumField()
	body, _ := ioutil.ReadAll(r.Body)
	var reqBody map[string]string

	err := json.Unmarshal(body, &reqBody)

	log.Println(reqBody)

	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i < numField; i++ {
		field := reflect.TypeOf(model).Field(i).Name
		tagValidate := reflect.TypeOf(model).Field(i).Tag.Get("validate")
		_, isExist := reqBody[field]

		log.Println(isExist)

		if tagValidate == "required" {
			if !isExist {
				return false
			}
		}

	}

	return true
}
