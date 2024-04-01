package mock_data

import (
	"context"
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
)

type MockServiceClientC interface {
	CreateComment(ctx context.Context, c *pbc.Comment) (*pbc.Comment, error)
	GetComment(ctx context.Context, c *pbc.CommentId) (*pbc.Comment, error)
	GetAllComment(ctx context.Context, c *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error)
	GetCommentsByPostId(ctx context.Context, c *pbc.PostId) (*pbc.GetAllCommentsResponse, error)
	GetCommentsByUserId(ctx context.Context, c *pbc.PostId) (*pbc.GetAllCommentsResponse, error)
	UpdateComment(ctx context.Context, c *pbc.Comment) (*pbc.Comment, error)
	DeleteComment(ctx context.Context, c *pbc.CommentId) (*pbc.DeleteResponse, error)
}

type mockServiceClientC struct{}

func NewMockServiceClientC() MockServiceClientC {
	return &mockServiceClientC{}
}

func (c *mockServiceClientC) CreateComment(ctx context.Context, in *pbc.Comment) (*pbc.Comment, error) {
	resp := &pbc.Comment{
		Id:       "testid",
		PostId:   "testpostid",
		UserId:   "testuserid",
		Content:  "testcontent",
		Likes:    0,
		Dislikes: 0,
	}
	return resp, nil
}

func (c *mockServiceClientC) GetComment(ctx context.Context, in *pbc.CommentId) (*pbc.Comment, error) {
	resp := &pbc.Comment{
		Id:       "testid",
		PostId:   "testpostid",
		UserId:   "testuserid",
		Content:  "testcontent",
		Likes:    0,
		Dislikes: 0,
	}
	return resp, nil
}

func (c *mockServiceClientC) GetAllComment(ctx context.Context, in *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error) {
	comments := []*pbc.Comment{
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
	}
	return &pbc.GetAllCommentsResponse{
		Comments: comments,
	}, nil
}

func (c *mockServiceClientC) GetCommentsByPostId(ctx context.Context, p *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	comments := []*pbc.Comment{
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
	}
	return &pbc.GetAllCommentsResponse{
		Comments: comments,
	}, nil
}

func (c *mockServiceClientC) GetCommentsByUserId(ctx context.Context, p *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	comments := []*pbc.Comment{
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
		{
			Id:       "testid",
			PostId:   "testpostid",
			UserId:   "testuserid",
			Content:  "testcontent",
			Likes:    0,
			Dislikes: 0,
		},
	}
	return &pbc.GetAllCommentsResponse{
		Comments: comments,
	}, nil
}

func (c *mockServiceClientC) UpdateComment(ctx context.Context, in *pbc.Comment) (*pbc.Comment, error) {
	resp := &pbc.Comment{
		Id:       "testid",
		PostId:   "testpostid",
		UserId:   "testuserid",
		Content:  "testcontent",
		Likes:    0,
		Dislikes: 0,
	}
	return resp, nil
}

func (c *mockServiceClientC) DeleteComment(ctx context.Context, in *pbc.CommentId) (*pbc.DeleteResponse, error) {
	resp := &pbc.DeleteResponse{
		Message: "success",
	}
	return resp, nil
}
