package model

import "github.com/jinzhu/gorm"

// Review of a user
type Review struct {
	gorm.Model
	UUID       string `json:"uuid"`
	ReviewerID uint
	RevieweeID uint
	Reviewer   User   `json:"reviewer"`
	Reviewee   User   `json:"reviewee"`
	Cards      []Card `json:"cards" gorm:"many2many:review_cards;"`
	Remark     string `json:"remark"`
	Completed  bool   `json:"completed"`
}
