package main

import (
	"fmt"
	"net/http"
	"portofolio/config"
	"portofolio/router"
)

func main() {
	router.InitRoute()
	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", config.Configuration().Server.Port)

	server.ListenAndServe()
}
