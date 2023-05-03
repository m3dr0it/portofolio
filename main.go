package main

import (
	"fmt"
	"log"
	"net/http"
	"portofolio/config"
	"portofolio/controller"
	"portofolio/db"
	"portofolio/middleware"
)

func main() {
	mux := http.NewServeMux()
	middleware.PrepareHeaderResponse(mux)
	mux.HandleFunc("/api/v1/whoami", controller.WhoAmi)
	mux.HandleFunc("/api/v1/login", controller.Login)
	mux.Handle("/api/v1/article", middleware.ValidateToken(controller.ArticlesController()))
	port := config.Configuration().Server.Port
	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: mux,
	}

	_, err := db.MongoDbClient()

	if err != nil {
		log.Fatal("mongo not connected")
	}

	server.ListenAndServe()
}
