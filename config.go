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

func WriteConfig(config *Config) error {
	databaseFile := databaseFile(DatabaseFile)
	return write(databaseFile, config)
}

func ReadConfig() (*Config, error) {
	databaseFile := databaseFile(DatabaseFile)
	return read(databaseFile)
}

func databaseFile(filename string) string {
	homePath, err := os.UserHomeDir()
	check(err)
	return fmt.Sprintf("%s/%s", homePath, filename)
}

func write(filename string, config *Config) error {
	configInJson, err := json.Marshal(config)
	check(err)

	return os.WriteFile(filename, configInJson, DefaultFilePermissions)
}

func read(filename string) (*Config, error) {
	configInJson, err := os.ReadFile(filename)
	check(err)

	var config = &Config{}
	err = json.Unmarshal(configInJson, &config)
	check(err)

	return config, nil
}

func clear(filename string) error {
	return os.Remove(filename)
}
