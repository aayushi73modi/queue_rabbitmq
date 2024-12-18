package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	// Declare a queue
	q, err := ch.QueueDeclare(
		"",    // Auto-generated queue name
		false, // Durable
		false, // Delete when unused
		true,  // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Bind the queue to the routing key 'logs.error'
	err = ch.QueueBind(
		q.Name,       // Queue name
		"logs.error", // Binding key
		"topic_logs", // Exchange name
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatal("Failed to bind a queue:", err)
	}

	// Start receiving messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	fmt.Println("Consumer 4 waiting for error messages...")
	for msg := range msgs {
		fmt.Printf("Consumer 4 received: %s\n", msg.Body)
	}
}
