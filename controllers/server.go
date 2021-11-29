package controllers

import (
	"encoding/json"
	"net/http"
	"time"

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

type Log struct {
	Act      string
	Stt      int
	CreateAt time.Time
}

func CreateLog(Act string, Stt int) Log {
	var log Log
	log.Act = Act
	log.Stt = Stt
	log.CreateAt = time.Now()
	return log
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
		log := CreateLog("C", http.StatusBadRequest)
		jsLog, _ := json.Marshal(log)
		loggers.Publish(string(jsLog))
		return
	}

	note := models.Note{User: input.User, Content: input.Content}
	models.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
	log := CreateLog("C", http.StatusOK)
	jsLog, _ := json.Marshal(log)
	loggers.Publish(string(jsLog))
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.First(&note, c.Param("ID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		log := CreateLog("U", http.StatusBadRequest)
		jsLog, _ := json.Marshal(log)
		loggers.Publish(string(jsLog))
		return
	}

	var input UpdateNoteInput
	if err := c.BindJSON((&input)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log := CreateLog("U", http.StatusBadRequest)
		jsLog, _ := json.Marshal(log)
		loggers.Publish(string(jsLog))
		return
	}

	models.DB.Model(&note).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": note})
	log := CreateLog("U", http.StatusOK)
	jsLog, _ := json.Marshal(log)
	loggers.Publish(string(jsLog))
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.First(&note, c.Param("ID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		log := CreateLog("D", http.StatusBadRequest)
		jsLog, _ := json.Marshal(log)
		loggers.Publish(string(jsLog))
		return
	}

	models.DB.Delete(&note)

	c.JSON(http.StatusOK, gin.H{"data": true})
	log := CreateLog("D", http.StatusOK)
	jsLog, _ := json.Marshal(log)
	loggers.Publish(string(jsLog))
}
