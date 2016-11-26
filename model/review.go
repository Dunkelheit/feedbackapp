package model

// Review of a user
type Review struct {
	ID        ID     `json:"id" db:"user_id, primarykey, autoincrement"`
	UUID      string `json:"uuid"`
	Reviewer  User   `json:"reviewer"`
	Reviewee  User   `json:"reviewee"`
	Cards     []Card `json:"cards"`
	Remark    string `json:"remark"`
	Completed bool   `json:"completed"`
}
