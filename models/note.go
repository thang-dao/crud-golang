package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	User    string `json:"user"`
	Content string `json:"content"`
}
