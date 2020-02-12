package main

import (
	"log"

	"github.com/mpuzanov/otus-go/calendar/internal/apiserver"
	flag "github.com/spf13/pflag"
)

var configFile string

func init() {
	flag.StringVarP(&configFile, "config", "c", "configs/config-dev.yml", "path config file")
	flag.Parse()
}

func main() {

	cfg, err := apiserver.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", configFile, err)
	}

	if err := apiserver.Start(cfg); err != nil {
		log.Fatal(err)
	}

}

// ./calendar --config=configs/config-dev.yml
// ./calendar --config=configs/config-prod.yml
// PORT=8091 ./calendar --config=configs/config-dev.yml
// curl -i localhost:8091
// curl -i localhost:8091/hello/Mikhail
