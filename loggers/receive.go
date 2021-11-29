package loggers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func Receive() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		panic("Failed to connect to connect RabbitMQ!")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		os.Getenv("QUEUE_NAME"), // name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		panic("Failed to delare a channel")
	}

	msgs, err := CHAN.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic("Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
