package messagebroker

import (
	// "context"
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"

	// transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	authLogService "github.com/alwismt/selectify/internal/adminApp/domains/shared/services"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumerListner() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println("Error: ", err)
		// return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"selectify_auth", // queue
		"",               // consumer
		true,             // autoAck
		false,            // exclusive
		false,            // noLocal
		false,            // noWait
		nil,              // arguments
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	forever := make(chan bool)
	go func() {
		eventData := new(transferobjects.EventAuthDTO)
		for d := range msgs {
			err := json.Unmarshal(d.Body, eventData)
			if err != nil {
				fmt.Println("Error decoding message:", err)
				continue
			}
			// fmt.Println("Received an Request")
			if eventData.Type == "login" {
				fmt.Println("Received a Login Request")
				authLog := authLogService.NewAuthLogService(nil)
				err := authLog.NewAuthEntry(context.Background(), eventData)
				if err != nil {
					fmt.Println(err)
				}
			}
			if eventData.Type == "logout" {
				fmt.Println("Received a Logout Request")
			}
			if eventData.Type == "logger" {
				fmt.Println("Received a Logger Request")
			}
			// fmt.Println("Received a message: ", eventData)
		}
	}()
	<-forever
	// return nil
}
