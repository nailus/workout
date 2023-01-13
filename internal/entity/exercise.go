package entity

type Exercise struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}