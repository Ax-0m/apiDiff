package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Project   string   `json:"project"`
	BaseURL   string   `json:"base_url"`
	Endpoints []string `json:"endpoints"`
}

func CreateConfig(projectName string) error {
	if _, err := os.Stat("apidiff.config.json"); err == nil {
		fmt.Println("apidiff.config.json already exists, skipping init")
		return nil
	}

	config := Config{
		Project:   projectName,
		BaseURL:   "https://api.example.com",
		Endpoints: []string{},
	}

	data, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		return fmt.Errorf("could not create config: %w", err)
	}

	err = os.WriteFile("apidiff.config.json", data, 0644)

	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	fmt.Println("created apidiff.config.json for project: ", projectName)
	return nil
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("apidiff.config.json")
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("could nor parse config file: %w", err)
	}

	return &config, nil
}
