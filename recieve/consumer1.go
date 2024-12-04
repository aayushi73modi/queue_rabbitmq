// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"

// 	"github.com/streadway/amqp"
// )

// func main() {
// 	// Connect to RabbitMQ server
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
// 	if err != nil {
// 		log.Fatal("Failed to connect to RabbitMQ:", err)
// 	}
// 	defer conn.Close()

// 	// Open a channel
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatal("Failed to open a channel:", err)
// 	}
// 	defer ch.Close()

// 	// Declare the queues
// 	queueName1 := "queue_task_1"
// 	queueName2 := "queue_task_2"

// 	_, err = ch.QueueDeclare(
// 		queueName1, // Queue name
// 		true,       // Durable
// 		false,      // Delete when unused
// 		false,      // Exclusive
// 		false,      // No-wait
// 		nil,        // Arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to declare queue %s: %v", queueName1, err)
// 	}

// 	_, err = ch.QueueDeclare(
// 		queueName2, // Queue name
// 		true,       // Durable
// 		false,      // Delete when unused
// 		false,      // Exclusive
// 		false,      // No-wait
// 		nil,        // Arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to declare queue %s: %v", queueName2, err)
// 	}
// 	// Use a WaitGroup to wait for both consumers
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		consumeMessages(ch, queueName1)
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		consumeMessages(ch, queueName2)
// 	}()
// 	wg.Wait()
// }

// // consumeMessages handles consuming messages from a single queue

// func consumeMessages(ch *amqp.Channel, queueName string) {
// 	msgs, err := ch.Consume(
// 		queueName, // Queue name
// 		"",        // Consumer tag
// 		true,      // Auto-acknowledge
// 		false,     // Exclusive
// 		false,     // No-local
// 		false,     // No-wait
// 		nil,       // Arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to register consumer for queue %s: %v", queueName, err)
// 	}
// 	for msg := range msgs {
// 		fmt.Printf("Received from %s: %s\n", queueName, msg.Body)
// 	}
// }

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

	// Declare a queue (must match the producer's queue declaration)
	q, err := ch.QueueDeclare(
		"queue_task_1", // Queue name
		true,           // Durable
		false,          // Delete when unused
		false,          // Exclusive
		false,          // No-wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}
	// Start receiving messages from the queue
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

	// Print the messages as they arrive
	for msg := range msgs {
		fmt.Printf("Received: %s\n", msg.Body)
	}
}
