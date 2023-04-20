package models

type Member struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Rank     string `json:"rank"`
	Vocation string `json:"vocation"`
	Level    int    `json:"level"`
	Joined   string `json:"joined"`
	Status   string `json:"status"`
}
