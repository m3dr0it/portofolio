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
	JwtExpired   int8   `json:"jwt_expired"`
	JwtSecretKey string `json:"jwt_secret_key"`
	Database     struct {
		Mongodb struct {
			Url        string `json:"url"`
			Port       int32  `json:"port"`
			Username   string `json:"username"`
			Password   string `json:"password"`
			Database   string `json:"database"`
			Collection string `json:"collection"`
		}
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
