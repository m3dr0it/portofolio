package router

import (
	"net/http"
	"portofolio/controller"
)

func InitApiRoute() {
	http.HandleFunc("/api/v1/login", controller.Login)
	http.HandleFunc("/api/v1/articles", controller.GetArticles().ServeHTTP)
}
