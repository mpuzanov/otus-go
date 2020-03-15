package scheduler

import (
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/scheduler"
	"github.com/mpuzanov/otus-go/calendar/internal/storage"
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
		Use:   "scheduler",
		Short: "Run scheduler",
		Run:   schedulerServerStart,
	}
)

func schedulerServerStart(cmd *cobra.Command, args []string) {

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
	mq, err := scheduler.New(cfg, logger, *db)
	if err != nil {
		log.Fatal(err)
	}
	if err := mq.Publish(); err != nil {
		log.Fatal(err)
	}
}
