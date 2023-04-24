package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var amqpChan *amqp.Channel

func RabbitMQConnection() error {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	amqpChan = ch
	return nil
}
