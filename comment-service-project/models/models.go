package models

type CreateComment struct {
    Id        string `json:"id"`
    PostId    string `json:"post_id"`
    UserId    string `json:"user_id"`
    Content   string `json:"content"`
    Likes     int64  `json:"likes"`
    Dislikes  int64  `json:"dislikes"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at"`
}
