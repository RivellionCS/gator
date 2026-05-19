package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	config_path, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("Error getting file path: %v\n", err)
	}

	data, err := os.ReadFile(config_path)
	if err != nil {
		fmt.Printf("Error reading file path: %v\n", err)
	}

	new_config := Config{}

	err = json.Unmarshal(data, &new_config)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
	}

	return new_config
}

func getConfigFilePath() (string, error) {
	config_file_path, err := os.UserHomeDir()
	config_file_path += "/.gatorconfig"
	return config_file_path, err
}

func (c Config) SetUser(username string) {
	c.CurrentUserName = username
	write(c)
}
