package model

import "github.com/jinzhu/gorm"

// User is somebody like you or me
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
	Avatar   string `json:"avatar"`
}
