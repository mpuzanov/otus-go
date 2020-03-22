package mq

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cenk/backoff"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// MQ структура процесса
type MQ struct {
	conf   *config.Config
	logger *zap.Logger
	conn   *amqp.Connection
	ch     *amqp.Channel
	q      amqp.Queue
	Done   chan error
}

// NewMQ создание процесса
func NewMQ(conf *config.Config, logger *zap.Logger) *MQ {

	return &MQ{conf: conf, logger: logger, Done: make(chan error)}
}

// Connect .
func (c *MQ) Connect() error {
	var err error
	c.conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		c.conf.Queue.User,
		c.conf.Queue.Password,
		c.conf.Queue.Host,
		c.conf.Queue.Port))
	if err != nil {
		return err
	}
	go func() {
		log.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		// Понимаем, что канал сообщений закрыт, надо пересоздать соединение.
		c.Done <- errors.New("Channel Closed")
	}()

	c.ch, err = c.conn.Channel()
	if err != nil {
		c.logger.Error("Failed to open a channel", zap.Error(err))
		return err
	}

	if err = c.ch.ExchangeDeclare(
		c.conf.Queue.ExchangeName,
		c.conf.Queue.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		c.logger.Error("Failed to declare an exchange", zap.Error(err))
		return err
	}

	c.q, err = c.ch.QueueDeclare(
		c.conf.Queue.QName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}

	c.logger.Info("Success connected to MQ")
	return nil
}

// ReConnect .
func (c *MQ) ReConnect() error {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = 5 * time.Minute
	be.InitialInterval = 1 * time.Second
	be.Multiplier = 2
	be.MaxInterval = 15 * time.Second

	b := backoff.WithContext(be, context.Background())
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return fmt.Errorf("stop reconnecting")
		}

		select {
		case <-time.After(d):
			if err := c.Connect(); err != nil {
				c.logger.Error("could not connect in reconnect call", zap.Error(err))
				continue
			}
			c.logger.Info("Reconnected... possibly")
			return nil
		}
	}
}

// AnnounceQueue Задекларировать очередь, которую будем слушать.
func (c *MQ) AnnounceQueue() (<-chan amqp.Delivery, error) {

	// Создаём биндинг (правило маршрутизации).
	if err := c.ch.QueueBind(
		c.q.Name,
		c.conf.Queue.BindingKey,
		c.conf.Queue.ExchangeName,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	msgs, err := c.ch.Consume(
		c.q.Name,
		c.conf.Queue.ConsumerTag,
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	return msgs, nil
}

// Publish отправка в очередь
func (c *MQ) Publish(body []byte) error {

	err := c.ch.Publish(
		c.conf.Queue.ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		c.logger.Error("Failed to publish a message", zap.Error(err))
		return err
	}
	c.logger.Debug("Sent", zap.String("Event", string(body)))
	return nil
}

// Close закрываем очередь
func (c *MQ) Close() (err error) {
	c.logger.Info("Close MQ connect...")
	err = c.ch.Close()
	if err != nil {
		_ = c.conn.Close()
		return err
	}
	err = c.conn.Close()
	if err == nil {
		c.logger.Info("Success close MQ connect")
	}
	return err
}
