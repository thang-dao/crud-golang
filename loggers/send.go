package loggers

import (
	"log"

	"github.com/streadway/amqp"
)

var QUEUE_NAME string
var CONN *amqp.Connection
var CHAN *amqp.Channel

func ConnectRabbitMQ(amqp_server_url string, queue_name string) {
	conn, err := amqp.Dial(amqp_server_url)
	if err != nil {
		panic("Failed to connect to connect RabbitMQ!")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		queue_name, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		panic("Failed to delare a channel")
	}

	QUEUE_NAME = q.Name
	CONN = conn
	CHAN = ch
}

func PublishMessage(body []byte) {
	err := CHAN.Publish(
		"",         // exchange
		QUEUE_NAME, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		panic("Failed to publish a message")
	}
	log.Printf(" [x] Sent %s\n", body)
}
