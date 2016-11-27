package model

// User is somebody like you or me
type User struct {
	ID       ID     `json:"id" orm:"auto;column(id)"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
	Avatar   string `json:"avatar" orm:"null"`
}
