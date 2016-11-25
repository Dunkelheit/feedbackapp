package model

// Review of a user
type Review struct {
	ID        ID `db:"user_id, primarykey, autoincrement"`
	UUID      string
	Reviewer  User
	Reviewee  User
	Cards     []Card
	Remark    string
	Completed bool
}
