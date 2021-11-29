package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thang-dao/crud-golang/loggers"
	"github.com/thang-dao/crud-golang/models"
)

type CreateNoteInput struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateNoteInput struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

func FindNoteByID(c *gin.Context) {
	var note models.Note

	if err := models.DB.First(&note, c.Param("ID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func FindNotes(c *gin.Context) {
	var notes []models.Note
	models.DB.Find(&notes)

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func CreateNote(c *gin.Context) {
	var input CreateNoteInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log := models.Log{Act: "C", Stt: http.StatusBadRequest}
		jsLog, _ := json.Marshal(log)
		loggers.PublishMessage(jsLog)
		return
	}

	note := models.Note{User: input.User, Content: input.Content}
	models.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
	log := models.Log{Act: "C", Stt: http.StatusOK}
	jsLog, _ := json.Marshal(log)
	loggers.PublishMessage(jsLog)
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.First(&note, c.Param("ID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		log := models.Log{Act: "U", Stt: http.StatusBadRequest}
		jsLog, _ := json.Marshal(log)
		loggers.PublishMessage(jsLog)
		return
	}

	var input UpdateNoteInput
	if err := c.BindJSON((&input)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log := models.Log{Act: "U", Stt: http.StatusBadRequest}
		jsLog, _ := json.Marshal(log)
		loggers.PublishMessage(jsLog)
		return
	}

	models.DB.Model(&note).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": note})
	log := models.Log{Act: "C", Stt: http.StatusOK}
	jsLog, _ := json.Marshal(log)
	loggers.PublishMessage(jsLog)
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.First(&note, c.Param("ID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		log := models.Log{Act: "D", Stt: http.StatusBadRequest}
		jsLog, _ := json.Marshal(log)
		loggers.PublishMessage(jsLog)
		return
	}

	models.DB.Delete(&note)

	c.JSON(http.StatusOK, gin.H{"data": true})
	log := models.Log{Act: "C", Stt: http.StatusOK}
	jsLog, _ := json.Marshal(log)
	loggers.PublishMessage(jsLog)
}
