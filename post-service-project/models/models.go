package models

type CreatePost struct {
	Id    string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Likes     int64  `json:"likes"`
	Dislikes  int64  `json:"dislikes"`
	Views     int64  `json:"views"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
}
