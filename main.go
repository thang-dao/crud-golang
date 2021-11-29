package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thang-dao/crud-golang/controllers"
	"github.com/thang-dao/crud-golang/loggers"
	"github.com/thang-dao/crud-golang/models"
	"github.com/thang-dao/crud-golang/pkg"
)

func main() {
	models.ConnectDatabaseNote()
	log.Print(pkg.AMQP_SERVER_URL)
	loggers.ConnectRabbitMQ(pkg.AMQP_SERVER_URL, pkg.QUEUE_NAME)
	router := gin.Default()

	router.GET("/notes", controllers.FindNotes)
	router.GET("/notes/:id", controllers.FindNoteByID)
	router.POST("/notes", controllers.CreateNote)
	router.PATCH("/notes/:id", controllers.UpdateNote)
	router.DELETE("/notes/:id", controllers.DeleteNote)

	router.Run(pkg.HOST + ":" + pkg.PORT)

	defer loggers.CONN.Close()
	defer loggers.CHAN.Close()

}
