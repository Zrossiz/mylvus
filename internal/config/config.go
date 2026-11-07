package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Zrossiz/mylvus/internal/helpers"
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
	SQLFilePath string `yaml:"sql_filepath"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Embedding struct {
	SearchMethod string `yaml:"search_method"`
}

func New() (*Config, error) {
	configsPath, err := helpers.WithBasePath("configs")
	if err != nil {
		return nil, fmt.Errorf("base path for public: %w", err)
	}

	env := os.Getenv("GOLANG_ENV")
	if env == "" {
		env = "development"
	} else {
		env = strings.ToLower(env)
	}

	filename := fmt.Sprintf("%s/%s.yml", configsPath, env)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		filename = configsPath + "/dev.yaml"
	}

	data, err := os.ReadFile(filename)
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
