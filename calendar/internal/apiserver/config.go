package apiserver

// Config Структура файла с конфигурацией
type Config struct {
	BindAddr    string `yaml:"http_listen"`
	LogLevel    string `yaml:"log_level"`
	LogFile     string `yaml:"log_file"`
	DatabaseURL string `yaml:"database_url"`
}
