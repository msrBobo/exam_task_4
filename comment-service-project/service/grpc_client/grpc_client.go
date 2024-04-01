package grpc_client

import (
	"exam_task_4/comment-service-project/config"
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
	pbp "exam_task_4/comment-service-project/genproto/post-service"
	pbu "exam_task_4/comment-service-project/genproto/user-service"
	mock "exam_task_4/comment-service-project/mock_data"
	"fmt"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	PostService() pbp.PostServiceClient
	CommentService() pbc.CommentServiceClient
	MockService() mock.MockServiceClientC
}

type serviceManager struct {
	cfg            config.Config
	userService    pbu.UserServiceClient
	postService    pbp.PostServiceClient
	commentService pbc.CommentServiceClient
	mockService    mock.MockServiceClientC
}

func New(cfg config.Config) (IServiceManager, error) {
	userServiceConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.RPCPort, cfg.RPCPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port: %s", cfg.RPCPort, cfg.RPCPort)
	}

	return &serviceManager{
		cfg:         cfg,
		userService: pbu.NewUserServiceClient(userServiceConn),
		mockService: mock.NewMockServiceClientC(),
	}, nil
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}

func (s *serviceManager) MockService() mock.MockServiceClientC {
	return s.mockService
}
