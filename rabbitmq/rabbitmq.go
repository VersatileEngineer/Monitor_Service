package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct{}

func (rmq *RabbitMQClient) ReceiveAddresses(RabbitMQURL string) {
	// Receive designated addresses from RabbitMQ and start monitoring transactions
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial(RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queueName := "address_queue"
	q, err := ch.QueueDeclare(
		queueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // auto-acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process incoming messages
	for msg := range msgs {
		address := string(msg.Body)
		log.Printf("Received address: %s\n", address)

		// Perform your desired actions with the received address, such as monitoring transactions and storing data in the database

		// Acknowledge the message to RabbitMQ
		err := msg.Ack(false)
		if err != nil {
			log.Printf("Failed to acknowledge message: %v\n", err)
			// You can choose to handle the error condition accordingly, such as retrying or logging the error.
		}
	}

	// Handle connection close gracefully
	<-conn.NotifyClose(make(chan *amqp.Error))
	log.Println("Connection to RabbitMQ closed")
}
