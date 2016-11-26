package model

// Card represents a single card
type Card struct {
	ID       ID           `json:"id" db:"user_id, primarykey, autoincrement"`
	Title    string       `json:"title"`
	Category CardCategory `json:"category"`
}
