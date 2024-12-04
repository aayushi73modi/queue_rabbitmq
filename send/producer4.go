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

	// Declare the direct exchange
	exchangeName := "topic_logs"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}
	//log.Println(q)

	// Publish error logs
	message := "queue log from Producer 4"
	err = ch.Publish(
		exchangeName,     // exchange
		"logs.queue4key", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
	fmt.Println("Producer 4 sent:", message)
}
