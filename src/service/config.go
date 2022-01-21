package service

import (
	"encoding/json"
	"os"
)

type Config struct {
	Http     *ConfigHttp     `json:"http"`
	Database *ConfigDatabase `json:"database"`
}

type ConfigDatabase struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type ConfigHttp struct {
	Endpoint string `json:"endpoint"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func GetConfiguration(path string) (*Config, error) {
	configData, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
