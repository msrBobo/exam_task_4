package mock_data

import (
	"context"
	pbp "exam_task_4/api-gateway-project/genproto/post-service"
)

type MockServiceClientP interface {
	CreatePost(ctx context.Context, p *pbp.Post) (*pbp.Post, error)
	GetPost(ctx context.Context, p *pbp.PostId) (*pbp.Post, error)
	GetAllPost(ctx context.Context, p *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error)
	GetPostsByUserId(ctx context.Context,  p *pbp.UserId) (*pbp.GetAllPostsResponse, error);
	UpdatePost(ctx context.Context, p *pbp.Post) (*pbp.Post, error)
	DeletePost(ctx context.Context, p *pbp.PostId) (*pbp.DeleteResponse, error)
}

type mockServiceClientPost struct{}

func NewMockServiceClientP() MockServiceClientP {
	return &mockServiceClientPost{}
}

func (c *mockServiceClientPost) CreatePost(ctx context.Context, p *pbp.Post) (*pbp.Post, error) {
	resp := &pbp.Post{
		Id:                   "testid",
		UserId:               "testuserid",
		Title:                "testtitle",
		Content:              "testcontent",
		Likes:                0,
		Dislikes:             0,
		Views:                0,
		Category:             "testcategory",
	}
	return resp, nil
}

func (c *mockServiceClientPost) GetPost(ctx context.Context, p *pbp.PostId) (*pbp.Post, error) {
	resp := &pbp.Post{
		Id:                   "testid",
		UserId:               "testuserid",
		Title:                "testtitle",
		Content:              "testcontent",
		Likes:                0,
		Dislikes:             0,
		Views:                0,
		Category:             "testcategory",
	}
	return resp, nil
}


func (c *mockServiceClientPost) GetAllPost(ctx context.Context, p *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error) {
	posts := []*pbp.Post{
		{
			Id:        "testid1",
			UserId:    "testuserid1",
			Title:     "testtitle1",
			Content:   "testcontent1",
			Likes:     0,
			Dislikes:  0,
			Views:     0,
			Category:  "testcategory1",
		},
		{
			Id:        "testid2",
			UserId:    "testuserid2",
			Title:     "testtitle2",
			Content:   "testcontent2",
			Likes:     0,
			Dislikes:  0,
			Views:     0,
			Category:  "testcategory2",
		},
	}
	return &pbp.GetAllPostsResponse{Posts: posts}, nil
}



func (c *mockServiceClientPost) GetPostsByUserId(ctx context.Context,  p *pbp.UserId) (*pbp.GetAllPostsResponse, error) {
	posts := []*pbp.Post{
		{
			Id:        "testid1",
			UserId:    "testuserid1",
			Title:     "testtitle1",
			Content:   "testcontent1",
			Likes:     0,
			Dislikes:  0,
			Views:     0,
			Category:  "testcategory1",
		},
		{
			Id:        "testid2",
			UserId:    "testuserid2",
			Title:     "testtitle2",
			Content:   "testcontent2",
			Likes:     0,
			Dislikes:  0,
			Views:     0,
			Category:  "testcategory2",
		},
	}
	return &pbp.GetAllPostsResponse{Posts: posts}, nil
}

func (c *mockServiceClientPost) UpdatePost(ctx context.Context, p *pbp.Post) (*pbp.Post, error) {
	resp := &pbp.Post{
		Id:                   "testid",
		UserId:               "testuserid",
		Title:                "testtitle",
		Content:              "testcontent",
		Likes:                0,
		Dislikes:             0,
		Views:                0,
		Category:             "testcategory",
	}
	return resp, nil
}

func (c *mockServiceClientPost) DeletePost(ctx context.Context, p *pbp.PostId) (*pbp.DeleteResponse, error) {
	resp := &pbp.DeleteResponse{
		Message: "success",
	}
	return resp, nil
}
