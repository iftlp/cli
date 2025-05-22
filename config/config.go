package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Theme string `json:"theme"`
}

func LoadConfig() (*Config, error) {
	configPath := "shizima.config"

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
