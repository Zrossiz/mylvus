package config

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

func New() *Config {
	var cfg Config

	// TODO: парсить yaml

	return &cfg
}
