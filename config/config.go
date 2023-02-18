package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type _Configuration struct {
	Server struct {
		Port int `json:port`
	}
}

var sharedConf *_Configuration

func init() {
	basePath, err := os.Getwd()

	log.Println(basePath)

	if err != nil {
		log.Println(err.Error())
		return
	}

	readConf, err := ioutil.ReadFile(filepath.Join(basePath, "config", "config.json"))

	if err != nil {
		log.Println(err.Error())
		return
	}

	sharedConf = new(_Configuration)

	errRead := json.Unmarshal(readConf, &sharedConf)

	if errRead != nil {
		log.Println(errRead.Error())
		return
	}

}

func Configuration() _Configuration {
	return *sharedConf
}
