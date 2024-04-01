package repo

import (
	pb "exam_task_4/user-service-project/genproto/user-service"
)

type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	GetUser(*pb.UserId) (*pb.User, error)
	GetAllUser(*pb.GetAllUsersRequest) (*pb.GetAllUsersResponse,error)
	UpdateUser(*pb.User) (*pb.User, error)
	DeleteUser(*pb.UserId) (*pb.DeleteResponse, error)
}
