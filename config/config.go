package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BindPort string `json:"bindport"`
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
