package scheduler

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/interfaces"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// Scheduler структура процесса
type Scheduler struct {
	conf   *config.Config
	conn   *amqp.Connection
	logger *zap.Logger
	db     interfaces.EventStorage
}

// New создание процесса
func New(conf *config.Config, logger *zap.Logger, db interfaces.EventStorage) (*Scheduler, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Queue.User, conf.Queue.Password, conf.Queue.Host, conf.Queue.Port))
	if err != nil {
		return nil, err
	}

	return &Scheduler{conf: conf, conn: conn, logger: logger, db: db}, nil
}

// Publish .
func (s *Scheduler) Publish() error {
	ch, err := s.conn.Channel()
	if err != nil {
		s.logger.Error("Failed to open a channel", zap.Error(err))
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		s.conf.Queue.Exchange, // name
		"fanout",              // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		s.logger.Error("Failed to declare an exchange", zap.Error(err))
		return err
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	tick := time.NewTicker(5 * time.Second)
loop:
	for {
		select {
		case <-stopChan:
			ch.Close()
			s.conn.Close()
			s.logger.Info("Exit the Scheduler")
			break loop
		case <-tick.C:
			// выбираем сообытия для отправки
			events, err := s.db.GetUserEvents("User1")
			if err != nil {
				s.logger.Error("Failed get events", zap.Error(err))
				break
			}
			s.logger.Debug("selected to send", zap.Int("count", len(events)))
			for _, event := range events {
				body, _ := json.Marshal(event)
				err = ch.Publish(
					s.conf.Queue.Exchange,
					"",
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        []byte(body),
					})
				if err != nil {
					s.logger.Error("Failed to publish a message", zap.Error(err))
					return err
				}
				s.logger.Info("Sent", zap.String("Event", string(body)))
			}
		}
	}
	return nil
}
