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
	Prom     PromConf       `yaml:"prometheus" mapstructure:"prometheus"`
}

// DBConf стуктура для настройки работы с базой данных
type DBConf struct {
	Name     string `yaml:"name" mapstructure:"name"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     string `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Database string `yaml:"database" mapstructure:"database"`
	SSL      string `yaml:"ssl" mapstructure:"ssl"`
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

// PromConf .
type PromConf struct {
	GRPCAddr   string `yaml:"grpc_addr" mapstructure:"grpc_addr"`
	SenderAddr string `yaml:"sender_addr" mapstructure:"sender_addr"`
}

// LoadConfig Загрузка конфигурации из файла
func LoadConfig(filePath string) (*Config, error) {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("log.level", "info")
	viper.SetDefault("http_listen", "localhost:8888")
	viper.SetDefault("grpc_listen", "localhost:50051")
	viper.SetDefault("prometheus.grpc_addr", "localhost:9091")
	viper.SetDefault("prometheus.sender_addr", "localhost:9092")
	viper.SetDefault("db.name", "postgres")
	viper.SetDefault("db.ssl", "disable")
	// QUEUE_HOST=192.168.56.103 DB_Database=pg_calendar_test

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
	if config.Log.Level == "debug" {
		log.Printf("config: %+v", config)
	}
	return &config, nil
}
