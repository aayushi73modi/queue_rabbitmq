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

	// declare  fanout exchange
	exchangeName := "fanout_exchange"
	//queueName := "queue_task_3"
	err = ch.ExchangeDeclare(
		exchangeName, // Exchange name
		"fanout",     // Exchange type
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Send a message to the queue
	body := "Hello to all queues via fanout exchange!"
	err = ch.Publish(
		exchangeName, // Fanout exchange name
		"",           // Routing key (ignored for fanout)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
	fmt.Println("Sent:", body)
}
