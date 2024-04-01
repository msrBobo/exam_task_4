package mock_data

import (
	"context"
	pbu "exam_task_4/user-service-project/genproto/user-service"
)

type MockServiceClientU interface {
	CreateUser(context.Context, *pbu.User) (*pbu.User, error)
	GetUser(context.Context, *pbu.UserId) (*pbu.User, error)
	UpdateUser(context.Context, *pbu.User) (*pbu.User, error)
	DeleteUser(context.Context, *pbu.UserId) (*pbu.DeleteResponse, error)
	GetAllUser(context.Context, *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error)
}

type mockServiceClientU struct{}

func NewMockServiceClientU() MockServiceClientU { return &mockServiceClientU{} }

func (c *mockServiceClientU) CreateUser(ctx context.Context, in *pbu.User) (*pbu.User, error) {
	return &pbu.User{
		Id:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Bio:       "testbio",
		Website:   "testwebsite",
		Username:  "testusername",
		Password:  "testpassword",
		Email:     "testemail",
	}, nil
}

func (c *mockServiceClientU) GetUser(ctx context.Context, in *pbu.UserId) (*pbu.User, error) {
	return &pbu.User{
		Id:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Bio:       "testbio",
		Website:   "testwebsite",
		Username:  "testusername",
		Password:  "testpassword",
		Email:     "testemail",
	}, nil
}

func (c *mockServiceClientU) GetAllUser(ctx context.Context, in *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	users := []*pbu.User{
		{
			Id:        "testid",
			FirstName: "testfirstname",
			LastName:  "testlastname",
			Bio:       "testbio",
			Website:   "testwebsite",
			Username:  "testusername",
			Password:  "testpassword",
			Email:     "testemail",
		},
		{
			Id:        "testid",
			FirstName: "testfirstname",
			LastName:  "testlastname",
			Bio:       "testbio",
			Website:   "testwebsite",
			Username:  "testusername",
			Password:  "testpassword",
			Email:     "testemail",
		},
	}
	return &pbu.GetAllUsersResponse{
		Users: users,
	}, nil
}

func (c *mockServiceClientU) UpdateUser(ctx context.Context, in *pbu.User) (*pbu.User, error) {
	return &pbu.User{
		Id:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Bio:       "testbio",
		Website:   "testwebsite",
		Username:  "testusername",
		Password:  "testpassword",
		Email:     "testemail",
	}, nil
}

func (c *mockServiceClientU) DeleteUser(ctx context.Context, in *pbu.UserId) (*pbu.DeleteResponse, error) {
	resp := &pbu.DeleteResponse{
		Message: "success",
	}
	return resp, nil
}
