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

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	// Declare the topic exchange
	err = ch.ExchangeDeclare(
		"topic_logs", // Exchange name
		"topic",      // Exchange type
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange:", err)
	}

	// Publishing messages with different routing keys
	messages := map[string]string{
		"logs.info":  "Info log message",
		"logs.debug": "Debug log message",
	}

	for key, message := range messages {
		err = ch.Publish(
			"topic_logs", // Exchange name
			key,          // Routing key
			false,        // Mandatory
			false,        // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			log.Fatal("Failed to publish a message:", err)
		}
		fmt.Printf("Sent: %s with key: %s\n", message, key)
	}
}
