package test

import (
	"log"
	"reflect"
	"testing"
)

type Test struct {
	Umur string `json:"name" validate:"required"`
}

func TestCekStruct(t *testing.T) {
	var test Test
	testAny(test)

}

func testAny(t any) {
	tagValidate := reflect.TypeOf(t).Field(0).Tag.Get("validate")
	fieldAlign := reflect.TypeOf(t).NumField()
	log.Println(tagValidate)
	log.Println(fieldAlign)

}
