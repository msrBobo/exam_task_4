package service

import (
	"context"
	pbp "exam_task_4/post-service-project/genproto/post-service"
	l "exam_task_4/post-service-project/pkg/logger"
	"exam_task_4/post-service-project/service/grpc_client"
	"exam_task_4/post-service-project/storage"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpc_client.IServiceManager
}

func NewPostService(client *mongo.Client, log l.Logger, grpcClient grpc_client.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStorageMongo(client),
		logger:  log,
		client:  grpcClient,
	}
}

func NewUserServicePostgres(db *sqlx.DB, log l.Logger, grpcClient grpc_client.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePostgres(db),
		logger:  log,
		client:  grpcClient,
	}
}

func (p *PostService) CreatePost(ctx context.Context, req *pbp.Post) (*pbp.Post, error) {
	post, err := p.storage.Post().CreatePost(req)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostService) GetPost(ctx context.Context, req *pbp.PostId) (*pbp.Post, error) {
	post, err := p.storage.Post().GetPost(req)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostService) GetAllPost(ctx context.Context, req *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error) {
	posts, err := p.storage.Post().GetAllPost(req)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostService) GetPostsByUserId(ctx context.Context, req *pbp.UserId) (*pbp.GetAllPostsResponse, error) {
	posts, err := p.storage.Post().GetPostsByUserId(req)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostService) UpdatePost(ctx context.Context, req *pbp.Post) (*pbp.Post, error) {
	post, err := p.storage.Post().UpdatePost(req)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostService) DeletePost(ctx context.Context, req *pbp.PostId) (*pbp.DeleteResponse, error) {
	resp, err := p.storage.Post().DeletePost(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
