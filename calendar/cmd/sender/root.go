package main

import (
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/mq"
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
		Use:   "sender",
		Short: "Run sender",
		Run:   senderServerStart,
	}
)

func senderServerStart(cmd *cobra.Command, args []string) {

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", cfgPath, err)
	}

	logger := logger.NewLogger(cfg.Log)

	mq := mq.NewMQ(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {

		err := mq.Connect()
		if err != nil {
			logger.Error("Connect MQ", zap.Error(err))
		}

		msgs, err := mq.AnnounceQueue()
		if err != nil {
			logger.Error("Announce Queue", zap.Error(err))
		}

		for {

			go func() {
				var countMsg int
				for d := range msgs {
					//logger.Debug("Send Event", zap.ByteString("Body", d.Body), zap.Int("countMsg", countMsg))
					log.Printf("Send Event: %s", d.Body)
					countMsg++
				}
			}()

			if <-mq.Done != nil {

				err = mq.ReConnect()
				if err != nil {
					logger.Error("Reconnecting Error: %s", zap.Error(err))
				} else {
					msgs, err = mq.AnnounceQueue()
					if err != nil {
						logger.Error("Failed AnnounceQueue", zap.Error(err))
						continue
					}
				}
			}
		}
	}()

	log.Printf("Waiting for logs. To exit press CTRL+C")
	<-forever
	mq.Close()

}
