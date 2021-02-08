package main

import (
	"encoding/json"
	"github.com/ariyn/Lcd/config"
	"io/ioutil"
	"os"
)

type Config struct {
	DB config.DB
}

func Load(path string) (conf Config) {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	return
}
