package repo

import (
	pbp "exam_task_4/post-service-project/genproto/post-service"
)

type PostStorageI interface {
	CreatePost(post *pbp.Post) (*pbp.Post, error)
	GetPost(*pbp.PostId) (*pbp.Post, error)
	GetPostsByUserId(*pbp.UserId) (*pbp.GetAllPostsResponse, error)
	GetAllPost(*pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error)
	UpdatePost(*pbp.Post) (*pbp.Post, error)
	DeletePost(*pbp.PostId) (*pbp.DeleteResponse, error)
}
