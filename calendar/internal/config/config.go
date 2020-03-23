package config

import (
	"log"
	"strings"

	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	"github.com/spf13/viper"
)

//Config Структура файла с конфигурацией
type Config struct {
	Log      logger.LogConf `yaml:"log" mapstructure:"log"`
	DB       DBConf         `yaml:"db" mapstructure:"db"`
	HTTPAddr string         `yaml:"http_listen" mapstructure:"http_listen"`
	GRPCAddr string         `yaml:"grpc_listen" mapstructure:"grpc_listen"`
	Queue    QueueConf      `yaml:"queue" mapstructure:"queue"`
}

// DBConf стуктура для настройки работы с базой данных
type DBConf struct {
	DbName      string `yaml:"db_name" mapstructure:"db_name"`
	DatabaseURL string `yaml:"database_url" mapstructure:"url"`
}

// QueueConf .
type QueueConf struct {
	Host         string `yaml:"host" mapstructure:"host"`
	Port         string `yaml:"port" mapstructure:"port"`
	User         string `yaml:"user" mapstructure:"user"`
	Password     string `yaml:"password" mapstructure:"password"`
	ExchangeName string `yaml:"exchange_name" mapstructure:"exchange_name"`
	ExchangeType string `yaml:"exchange_type" mapstructure:"exchange_type"`
	QName        string `yaml:"qname" mapstructure:"qname"`
	BindingKey   string `yaml:"binding_key" mapstructure:"binding_key"`
	ConsumerTag  string `yaml:"consumer_tag" mapstructure:"consumer_tag"`
}

// LoadConfig Загрузка конфигурации из файла
func LoadConfig(filePath string) (*Config, error) {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("log.level", "info")
	viper.SetDefault("http_listen", "0.0.0.0:8090")
	viper.SetDefault("grpc_listen", "0.0.0.0:50051")
	viper.SetDefault("db.url", "postgres://postgres:12345@localhost:5432/pg_calendar_test?sslmode=disable")

	if filePath != "" {
		log.Printf("Parsing config: %s\n", filePath)
		viper.SetConfigFile(filePath)
		viper.SetConfigType("yaml")
		//log.Println(viper.ConfigFileUsed())
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Config file is not specified.")
	}
	//log.Println(viper.AllSettings())

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	//log.Println(config)
	return &config, nil
}
