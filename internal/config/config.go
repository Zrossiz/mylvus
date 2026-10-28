package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        App        `yaml:"app"`
	Server     Server     `yaml:"server"`
	Repository Repository `yaml:"repository"`
	Embedding  Embedding  `yaml:"embedding"`
}

type App struct {
	ENV         string `yaml:"env"`
	LoggerLevel string `yaml:"logger_level"`
}

type Repository struct {
	JSONFilePath string `yaml:"json_file_path"`
}

type Server struct {
	Address string
	Port    string
}

type Embedding struct {
	Length       int    `yaml:"length"`
	SearchMethod string `yaml:"search_method"`
}

func New() (*Config, error) {
	data, err := os.ReadFile("/Users/zrossiz/Desktop/mylvus/config/dev.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return &cfg, nil
}
