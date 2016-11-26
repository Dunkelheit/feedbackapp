package model

// User is somebody like you or me
type User struct {
	ID       ID     `json:"id" db:"user_id, primarykey, autoincrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle" db:"job_title"`
	Avatar   string `json:"avatar"`
}
