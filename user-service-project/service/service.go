package service

import (
	"context"
	pbu "exam_task_4/user-service-project/genproto/user-service"
	l "exam_task_4/user-service-project/pkg/logger"
	"exam_task_4/user-service-project/service/grpc_client"
	"exam_task_4/user-service-project/storage"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpc_client.IServiceManager
}

func NewUserServiceMongo(client *mongo.Client, log l.Logger, grpcClient grpc_client.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStorageMongo(client),
		logger:  log,
		client:  grpcClient,
	}
}

func NewUserServicePostgres(db *sqlx.DB, log l.Logger, grpcClient grpc_client.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePostgres(db),
		logger:  log,
		client:  grpcClient,
	}
}


func (u *UserService) CreateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := u.storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pbu.UserId) (*pbu.User, error) {
	user, err := u.storage.User().GetUser(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	users, err := u.storage.User().GetAllUser(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := u.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pbu.UserId) (*pbu.DeleteResponse, error) {
	resp, err := u.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
