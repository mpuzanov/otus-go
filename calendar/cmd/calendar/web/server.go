package web

import (
	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/storage"
	"github.com/mpuzanov/otus-go/calendar/internal/web"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"

	"log"

	"github.com/spf13/cobra"
)

var cfgPath string

func init() {
	ServerCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
}

var (
	// ServerCmd .
	ServerCmd = &cobra.Command{
		Use:   "web_server",
		Short: "Run web server",
		Run:   webServerStart,
	}
)

func webServerStart(cmd *cobra.Command, args []string) {

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", cfgPath, err)
	}
	logger := logger.NewLogger(cfg.Log)
	// Init db
	db, err := storage.NewStorageDB(cfg)
	if err != nil {
		log.Fatalf("newStorageDB failed: %s", err)
	}
	calendar := calendar.NewCalendar(*db)

	if err := web.Start(cfg, logger, calendar); err != nil {
		log.Fatal(err)
	}
}
