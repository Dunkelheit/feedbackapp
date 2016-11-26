package model

// Review of a user
type Review struct {
	ID        ID      `json:"id" orm:"auto"`
	UUID      string  `json:"uuid"`
	Reviewer  *User   `json:"reviewer" orm:"rel(one)"`
	Reviewee  *User   `json:"reviewee" orm:"rel(one)"`
	Cards     []*Card `json:"cards" orm:"rel(m2m)"`
	Remark    string  `json:"remark"`
	Completed bool    `json:"completed"`
}
