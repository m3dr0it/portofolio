package controller

import (
	"fmt"
	"net/http"
)

func GetArticles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	})
}
