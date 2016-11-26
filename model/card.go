package model

// Card represents a single card
type Card struct {
	ID       ID           `json:"id" orm:"auto"`
	Title    string       `json:"title"`
	Category CardCategory `json:"category"`
}
