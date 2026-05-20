package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type state struct {
	config *Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	command_list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	c.command_list[cmd.name]
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	s.config.SetUser(cmd.arguments[2])
	fmt.Println("the user has been set")
	return nil
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
	config_file_path += "/" + configFileName
	return config_file_path, err
}

func (c Config) SetUser(username string) {
	c.CurrentUserName = username
	write(c)
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error could not convert to json: %v\n", err)
	}
	config_file_path, _ := os.UserHomeDir()
	config_file_path += "/" + configFileName
	err = os.WriteFile(config_file_path, jsonData, 0664)
	if err != nil {
		return fmt.Errorf("Error writing to the config file: %v\n", err)
	}
	return err
}