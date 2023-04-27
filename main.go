package main

import (
	"fmt"
	"net/http"
	"portofolio/config"
	"portofolio/controller"
	"portofolio/middleware"
)

func main() {
	mux := http.NewServeMux()
	middleware.PrepareHeaderResponse(mux)
	mux.HandleFunc("/api/v1/whoami", controller.WhoAmi)
	mux.HandleFunc("/api/v1/login", controller.Login)
	mux.Handle("/api/v1/article", middleware.ValidateToken(controller.GetArticles()))
	port := config.Configuration().Server.Port
	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: mux,
	}

	server.ListenAndServe()
}
