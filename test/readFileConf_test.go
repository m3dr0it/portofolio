package test

import (
	"log"
	"portofolio/config"
	"testing"
)

func TestReadFile(t *testing.T) {
	port := config.Configuration().Server.Port
	log.Println(port)
}
