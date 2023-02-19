package util

import (
	"log"
	"reflect"
)

func ValidatePayload(body map[string]string, model any) bool {
	numField := reflect.TypeOf(model).NumField()

	for i := 0; i < numField; i++ {
		field := reflect.TypeOf(model).Field(i).Tag.Get("json")
		tagValidate := reflect.TypeOf(model).Field(i).Tag.Get("validate")
		_, isExist := body[field]

		log.Println(isExist)

		if tagValidate == "required" {
			if !isExist {
				return false
			}
		}

	}

	return true
}
