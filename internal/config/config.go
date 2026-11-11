package config

import (
	"encoding/json"
	"log"
	"os"
)

type Cfg struct {
	Url Urls `json:"urls"`
}

type Urls struct {
	MainPage     string `json:"mainPage"`
	RaplacedPage string `json:"raplacedPage"`
}

func Config(path string) *Cfg {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("не удалось прочитать файл: %v", err)
	}

	var cfg Cfg
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("parsing err: %v", err)
	}

	return &cfg
}
