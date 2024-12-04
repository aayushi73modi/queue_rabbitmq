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
	log.Println(q)

	// Send a message to the queue
	body := "Hello task queue 1"
	err = ch.Publish(
		"",     // Default exchange
		q.Name, // Routing key (queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	log.Println(q.Name)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
	fmt.Println("Sent:", body)
}
