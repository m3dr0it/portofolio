package router

import (
	"net/http"
	"portofolio/controller"
)

func InitRoute() {
	http.HandleFunc("/api/v1/login", controller.Login)
}
