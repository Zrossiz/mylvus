package config

type Config struct {
	App        App        `yaml:"app"`
	Server     Server     `yaml:"server"`
	Repository Repository `yaml:"repository"`
}

type App struct {
	ENV string `yaml:"env"`
}

type Repository struct {
	JSONFilePath string `yaml:"json_file_path"`
}

type Server struct {
	Address string
	Port    string
}

func New() *Config {
	var cfg Config

	// TODO: парсить yaml

	return &cfg
}
