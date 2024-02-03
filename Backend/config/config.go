package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Server          ServerConfig          `json:"server"`
	Database        DatabaseConfig        `json:"database"`
	SpoonacularAPI  SpoonacularAPIConfig  `json:"spoonacular_api"`
}

type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type DatabaseConfig struct {
	ConnectionString string `json:"connection_string"`
	Name             string `json:"name"`
}

type SpoonacularAPIConfig struct {
	BaseURL string `json:"base_url"`
	APIKey  string `json:"api_key"`
}

func LoadConfig(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
