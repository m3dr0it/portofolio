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
	mux.Handle("/api/v1/article", middleware.ValidateToken(controller.GetArticles()))
	mux.HandleFunc("/api/v1/login", controller.Login)
	port := config.Configuration().Server.Port
	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: mux,
	}

	server.ListenAndServe()
}
