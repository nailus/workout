package entity

type Exercise struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	AuthorId  int `json:"author_id" db:"author_id"`
}