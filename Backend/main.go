package main;

import (
	"fmt"
	"github.com/professorabhay/config"
)

func main(){
	fmt.Println("Hello SavorSphere!");

	configPath := "config/config.json"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	fmt.Printf("Server Port: %d\n", cfg.Server.Port)
	fmt.Printf("Database Connection String: %s\n", cfg.Database.ConnectionString)
	fmt.Printf("Spoonacular API Key: %s\n", cfg.SpoonacularAPI.APIKey)
}