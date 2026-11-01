package config

import (
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
	BinaryFilePath string `yaml:"binary_file_path"`
	SQLFilePath    string `yaml:"sql_file_path"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Embedding struct {
	SearchMethod string `yaml:"search_method"`
}

func New() (*Config, error) {
	data, err := os.ReadFile("/Users/zrossiz/Desktop/vectorDB/mylvus/config/dev.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
