package service

import (
	"exam_task_4/api-gateway-project/config"
	pbc "exam_task_4/api-gateway-project/genproto/comment-service"
	pbp "exam_task_4/api-gateway-project/genproto/post-service"
	pbu "exam_task_4/api-gateway-project/genproto/user-service"
	mock "exam_task_4/api-gateway-project/mock_data"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	PostService() pbp.PostServiceClient
	CommentService() pbc.CommentServiceClient
	MockServiceU() mock.MockServiceClientU
	MockServiceP() mock.MockServiceClientP
	MockServiceC() mock.MockServiceClientC
}

type serviceManager struct {
	userService    pbu.UserServiceClient
	postService    pbp.PostServiceClient
	commentService pbc.CommentServiceClient
	mockServiceU    mock.MockServiceClientU
	mockServiceP    mock.MockServiceClientP
	mockServiceC    mock.MockServiceClientC
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

func (s *serviceManager) MockServiceU() mock.MockServiceClientU {
	return s.mockServiceU
}

func (s *serviceManager) MockServiceP() mock.MockServiceClientP {
	return s.mockServiceP
}

func (s *serviceManager) MockServiceC() mock.MockServiceClientC {
	return s.mockServiceC
}
func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CommentServiceHost, conf.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pbu.NewUserServiceClient(connUser),
		postService:    pbp.NewPostServiceClient(connPost),
		commentService: pbc.NewCommentServiceClient(connComment),
		mockServiceU:    mock.NewMockServiceClientU(),
		mockServiceP:    mock.NewMockServiceClientP(),
		mockServiceC:    mock.NewMockServiceClientC(),
	}
	return serviceManager, nil
}
