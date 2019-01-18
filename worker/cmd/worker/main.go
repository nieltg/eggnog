package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatalf("Unable to dial Message Queue: %s", err)
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("Unable to open a channel: %s", err)
	}

	queue, err := channel.QueueDeclare(
		"hello", // name
		true,    // durable
		true,    // delete when unused
		false,   //exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Unable to declare a queue: %s", err)
	}

	messages, err := channel.Consume(
		queue.Name, // queue name
		"",         // consumer name
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Unable to register a consumer: %s", err)
	}

	for message := range messages {
		log.Printf("Received a message: %s", message.Body)

		err := channel.Ack(message.DeliveryTag, false)
		if err != nil {
			log.Fatalf("Unable to send ACK: %s", err)
		}
	}
}
