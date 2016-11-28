package model

// User is somebody like you or me
type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
	Avatar   string `json:"avatar"`
}
