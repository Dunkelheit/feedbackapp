package model

// User is somebody like you or me
type User struct {
	ID       ID     `db:"user_id, primarykey, autoincrement"`
	Name     string ``
	Email    string ``
	JobTitle string `db:"job_title"`
	Avatar   string ``
}
