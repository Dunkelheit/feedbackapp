package model

// Card represents a single card
type Card struct {
	ID       ID `db:"user_id, primarykey, autoincrement"`
	Title    string
	Category CardCategory
}
