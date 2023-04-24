package queue

import (
	"context"
	"encoding/json"
	"fmt"

	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher interface {
	QueueEvent(event interface{}) error
}

type rabbitMQPublisher struct {
	// conn *amqp.Connection
	ch *amqp.Channel
}

func NewRabbitMQPublisher(ch *amqp.Channel) RabbitMQPublisher {
	if ch == nil {
		ch = amqpChan
	}
	return &rabbitMQPublisher{ch: ch}
}

func (r *rabbitMQPublisher) QueueEvent(event interface{}) error {
	var name string
	// var err error
	// var q amqp.Queue
	switch e := event.(type) {
	case *transferobjects.EventAuthDTO:
		name = e.Name
	case *transferobjects.QueueDTO:
		name = e.Name
	case *transferobjects.EmailDTO:
		name = e.Name
	default:
		return fmt.Errorf("unsupported event type: %T", event)

	}

	q, err := r.ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	// create a json message from the event
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// fmt.Println(" [x] Publishing event: ", string(body))
	err = r.ch.PublishWithContext(
		context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	return nil
}
