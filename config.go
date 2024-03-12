package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DatabaseFile           = ".snappass"
	DefaultFilePermissions = 0644
)

type Config struct {
	BaseUrl string `json:"base_url"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func databaseFile() string {
	homePath, err := os.UserHomeDir()
	check(err)
	return fmt.Sprintf("%s/%s", homePath, DatabaseFile)
}

func Write(config *Config) error {
	configInJson, err := json.Marshal(config)
	check(err)

	return os.WriteFile(databaseFile(), configInJson, DefaultFilePermissions)
}

func Read() (*Config, error) {
	configInJson, err := os.ReadFile(databaseFile())
	check(err)

	var config = &Config{}
	err = json.Unmarshal(configInJson, &config)
	check(err)

	return config, nil
}
