package model

// Card represents a single card
type Card struct {
	ID       ID           `json:"id" orm:"auto"`
	Title    string       `json:"title" binding:"required"`
	Category CardCategory `json:"category" binding:"required"`
}
