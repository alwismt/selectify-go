package messagebroker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQListener struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQListener(connString string, queueName string) (*RabbitMQListener, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQListener{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (l *RabbitMQListener) StartListening(handlerFunc func([]byte)) error {
	msgs, err := l.channel.Consume(
		l.queue.Name, // queue
		"",           // consumer
		false,        // autoAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			handlerFunc(d.Body)
			d.Ack(false)
		}
	}()

	return nil
}

func (l *RabbitMQListener) StopListening() error {
	if err := l.channel.Close(); err != nil {
		return err
	}

	if err := l.conn.Close(); err != nil {
		return err
	}

	return nil
}
