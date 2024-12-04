package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	// Declare the direct exchange
	exchangeName := "direct_logs"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange:", err)
	}

	// Publish info logs
	infoMessage := "queue log from Producer 2"
	err = ch.Publish(
		exchangeName, // exchange
		"info",       // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(infoMessage),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
	fmt.Println("Producer 2 sent:", infoMessage)

	// Publish warning logs
	warningMessage := "Warning log from Producer 2"
	err = ch.Publish(
		exchangeName, // exchange
		"queue2key",  // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(warningMessage),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
	fmt.Println("Producer 2 sent:", warningMessage)
}
