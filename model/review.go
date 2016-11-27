package model

// Review of a user
type Review struct {
	ID        ID      `json:"id" orm:"auto;column(id)"`
	UUID      string  `json:"uuid" orm:"column(uuid)"`
	Reviewer  *User   `json:"reviewer" orm:"rel(fk);on_delete(cascade)"`
	Reviewee  *User   `json:"reviewee" orm:"rel(fk);on_delete(cascade)"`
	Cards     []*Card `json:"cards" orm:"rel(m2m)"`
	Remark    string  `json:"remark"`
	Completed bool    `json:"completed"`
}
