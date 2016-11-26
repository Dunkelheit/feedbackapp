package model

// User is somebody like you or me
type User struct {
	ID       ID     `json:"id" orm:"auto"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
	Avatar   string `json:"avatar"`
}
