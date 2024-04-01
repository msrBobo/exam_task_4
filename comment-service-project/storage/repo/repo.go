package repo

import (
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
)

type CommentStorageI interface {
	CreateComment(*pbc.Comment) (*pbc.Comment, error)
	GetComment(*pbc.CommentId) (*pbc.Comment, error)
	GetAllComment(*pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error)
	GetCommentsByPostId(*pbc.PostId) (*pbc.GetAllCommentsResponse, error)
	GetCommentsByUserId(*pbc.PostId) (*pbc.GetAllCommentsResponse, error) 
	UpdateComment(*pbc.Comment) (*pbc.Comment, error)
	DeleteComment(*pbc.CommentId) (*pbc.DeleteResponse, error)
}
