package storage

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio		  string `json:"bio"`
	Website   string `json:"website"`
}

type Post struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Likes    int64  `json:"likes"`
	Dislikes int64  `json:"dislikes"`
	Views    int64  `json:"views"`
	Category string `json:"category"`
}

type Comment struct {
	Id       string `json:"id"`
	PostId   string `json:"post_id"`
	UserId   string `json:"user_id"`
	Content  string `json:"content"`
	Likes    int64  `json:"likes"`
	Dislikes int64  `json:"dislikes"`
}