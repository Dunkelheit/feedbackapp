package model

import "github.com/jinzhu/gorm"

// Card represents a single card
type Card struct {
	gorm.Model
	Title    string       `json:"title" binding:"required"`
	Category CardCategory `json:"category" binding:"exists"`
}
