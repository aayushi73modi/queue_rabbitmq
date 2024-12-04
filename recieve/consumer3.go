package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
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

	// Declare a queue (this is done by the consumer)
	name := "queue_task_1"
	q, err := ch.QueueDeclare(
		name,  // Queue name (can be unique for each consumer)
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		q.Name,            // Queue name
		"",                // Routing key (not used in fanout)
		"fanout_exchange", // Exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind queue:", err)
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

	// Handle messages
	for msg := range msgs {
		fmt.Printf("Received: %s\n", msg.Body)
	}
}
