package sender

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// Sender .
type Sender struct {
	conf   *config.Config
	conn   *amqp.Connection
	logger *zap.Logger
}

// NewSender .
func NewSender(conf *config.Config, logger *zap.Logger) (*Sender, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Queue.User, conf.Queue.Password, conf.Queue.Host, conf.Queue.Port))
	if err != nil {
		return nil, err
	}

	return &Sender{conf: conf, conn: conn, logger: logger}, nil
}

// ReadSendMessage .
func (s *Sender) ReadSendMessage() error {

	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
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
		return err
	}

	q, err := ch.QueueDeclare(
		s.conf.Queue.Name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,                // queue name
		"sending",             // routing key
		s.conf.Queue.Exchange, // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	msgsChan, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

loop:
	for {
		select {
		case <-stopChan:
			s.logger.Info("Exit Sender")
			ch.Close()
			s.conn.Close()
			break loop
		case msg, ok := <-msgsChan:
			if !ok {
				break loop
			}
			log.Println(string(msg.Body))
		}
	}

	return nil
}
