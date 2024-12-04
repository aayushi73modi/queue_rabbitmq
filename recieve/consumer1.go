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

	// Declare a queue
	queueName := "queue_task_1"
	q, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	// Process messages
	fmt.Println("Consumer 1 is waiting for messages...")
	for msg := range msgs {
		fmt.Printf("Consumer 1 received: %s\n", msg.Body)
	}
}
