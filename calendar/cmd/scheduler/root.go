package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/mq"
	"github.com/mpuzanov/otus-go/calendar/internal/storage"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	"go.uber.org/zap"

	"log"

	"github.com/spf13/cobra"
)

var cfgPath string

func init() {
	serverCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
}

var (
	serverCmd = &cobra.Command{
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
	// logger := logger.NewLogger(
	// 	logger.LogConf{
	// 		Level:      cfg.Log.Level,
	// 		File:       cfg.Log.File,
	// 		FormatJSON: cfg.Log.FormatJSON,
	// 	},
	// )
	logger := logger.NewLogger(cfg.Log)

	db, err := storage.NewStorageRemind(cfg)
	if err != nil {
		log.Fatalf("NewStorageRemind failed: %s", err)
	}

	mq := mq.NewMQ(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {

		err := mq.Connect()
		if err != nil {
			log.Fatal(err)
		}

		ticker := time.NewTicker(10 * time.Second)
		for {

			// выбираем сообытия для отправки
			events, err := db.GetUserEvents("User1")
			if err != nil {
				logger.Error("Failed get events", zap.Error(err))
				break
			}
			logger.Info("selected to send", zap.Int("count", len(events)))
			for _, event := range events {
				body, err := json.Marshal(event)
				if err != nil {
					logger.Error("Marshal event", zap.Error(err))
				}
				err = mq.Publish(body)
				if err != nil {
					logger.Error("Failed to publish a message", zap.Error(err))

					err := mq.ReConnect()
					if err != nil {
						logger.Error("Fail reconnect to MQ:", zap.Error(err))
					}
					continue
				}
			}
			<-ticker.C
		}
		ticker.Stop()
	}()

	<-stopChan
	mq.Close()
}
