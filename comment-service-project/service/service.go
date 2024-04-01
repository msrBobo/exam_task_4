package service

import (
	"context"
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
	l "exam_task_4/comment-service-project/pkg/logger"
	"exam_task_4/comment-service-project/service/grpc_client"
	"exam_task_4/comment-service-project/storage"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpc_client.IServiceManager
}

func NewCommentServiceMongo(client *mongo.Client, log l.Logger, grpcClient grpc_client.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStorageMongo(client),
		logger:  log,
		client:  grpcClient,
	}
}

func NewUserServicePostgres(db *sqlx.DB, log l.Logger, grpcClient grpc_client.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePostgres(db),
		logger:  log,
		client:  grpcClient,
	}
}

func (c *CommentService) CreateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	comment, err := c.storage.Comment().CreateComment(req)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentService) GetComment(ctx context.Context, req *pbc.CommentId) (*pbc.Comment, error) {
	comment, err := c.storage.Comment().GetComment(req)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentService) GetAllComment(ctx context.Context, req *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error) {
	comments, err := c.storage.Comment().GetAllComment(req)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentService) GetCommentsByPostId(ctx context.Context, req *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	comments, err := c.storage.Comment().GetCommentsByPostId(req)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentService) GetCommentsByUserId(ctx context.Context, req *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	comments, err := c.storage.Comment().GetCommentsByUserId(req)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentService) UpdateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	comment, err := c.storage.Comment().UpdateComment(req)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentService) DeleteComment(ctx context.Context, req *pbc.CommentId) (*pbc.DeleteResponse, error) {
	resp, err := c.storage.Comment().DeleteComment(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
