package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BindPort string `json:"bindport"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	SSLmode  string `json:"sslmode"`
}

func ReadConfig(path string) Config {
	var err error
	var cfg Config
	if content, err := ioutil.ReadFile(path); err == nil {
		if err = json.Unmarshal(content, &cfg); err == nil {
			log.Println(cfg)
		}
	}

	if err != nil {
		log.Printf("ReadJson err: %s", err.Error())
	}

	return cfg
}
