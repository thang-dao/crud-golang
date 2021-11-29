package pkg

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/thang-dao/crud-golang/tools"
)

var (
	HOST            string
	PORT            string
	AMQP_SERVER_URL string
	QUEUE_NAME      string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	HOST = tools.Getenv("HOST", "localhost")
	PORT = tools.Getenv("PORT", "8000")
	AMQP_SERVER_URL = tools.Getenv("AMQP_SERVER_URL", "waring")
	QUEUE_NAME = tools.Getenv("QUEUE_NAME", "")
}
