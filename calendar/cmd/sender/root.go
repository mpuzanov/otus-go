package main

import (
	"net/http"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/mq"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	"go.uber.org/zap"

	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// metric for prometheus
	msgCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "rmq_sent_messages",
		Help: "Count of messages sent in seconds",
	})
)

func senderServerStart(cmd *cobra.Command, args []string) {

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", cfgPath, err)
	}

	logger := logger.NewLogger(cfg.Log)

	prometheus.MustRegister(msgCounter)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("mq metrics initialisation...")
		log.Fatal(http.ListenAndServe(cfg.Prom.SenderAddr, nil))
	}()

	mq := mq.NewMQ(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	forever := make(chan bool)

	go func() {

		err := mq.Connect()
		if err != nil {
			log.Fatalf("Failed connect MQ: %s", err)
		}

		msgs, err := mq.AnnounceQueue()
		if err != nil {
			log.Fatalf("Announce Queue: %s", err)
		}

		for {

			go func() {
				var countMsg int
				for d := range msgs {
					//logger.Debug("Send Event", zap.ByteString("Body", d.Body), zap.Int("countMsg", countMsg))
					log.Printf("Send Event: %s, CountMsg: %d", d.Body, countMsg)
					countMsg++
					msgCounter.Inc()
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

}
