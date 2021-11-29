package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/thang-dao/crud-golang/models"
	"github.com/thang-dao/crud-golang/pkg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabaseLog() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("log.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Log{})

	DB := database
	return DB
}

func main() {
	DB := ConnectDatabaseLog()
	conn, err := amqp.Dial(pkg.AMQP_SERVER_URL)
	if err != nil {
		panic("Failed to connect to connect RabbitMQ!")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		pkg.QUEUE_NAME,  // queue
		"CRUD-CONSUMER", // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Print(err.Error())
		panic("Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var messLog models.Log
			err = json.Unmarshal(d.Body, &messLog)

			DB.Create(&messLog)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
