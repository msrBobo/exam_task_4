package mongosh

import (
	"context"
	"errors"
	pb "exam_task_4/user-service-project/genproto/user-service"
	m "exam_task_4/user-service-project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepoMongo(collection *mongo.Collection) *userRepo {
	return &userRepo{collection: collection}
}

func (u *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	user.CreatedAt = time.Now().String()
	CreateUser := m.CreateUser{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Bio:       user.Bio,
		Website:   user.Website,
	}
	_, err := u.collection.InsertOne(context.Background(), CreateUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) GetUser(id *pb.UserId) (user *pb.User, err error) {
	filter := bson.M{"id": id.UserId}
	err = u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) GetAllUser(req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	ctx := context.TODO()
	offset := (req.Page - 1) * req.Limit

	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(req.Limit))

	filter := bson.M{"deleted_at": nil}

	cursor, err := u.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*pb.User
	for cursor.Next(ctx) {
		var post pb.User
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		users = append(users, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetAllUsersResponse{Users: users}, nil
}

func (u *userRepo) UpdateUser(user *pb.User) (*pb.User, error) {
	filter := bson.M{"id": user.Id}
	user.UpdatedAt = time.Now().String()
	update := bson.M{
		"$set": bson.M{
			"username":   user.Username,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"bio":        user.Bio,
			"website":    user.Website,
		},
	}
	_, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	var updatedUser pb.User
	err = u.collection.FindOne(context.Background(), filter).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (u *userRepo) DeleteUser(id *pb.UserId) (*pb.DeleteResponse, error) {
	deletedAt := time.Now()
	filter := bson.M{"id": id.UserId, "deleted_at": nil}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}
	result, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("delete failed: post not found or already deleted")
	}
	resp := &pb.DeleteResponse{
		Message: "Deleted post",
	}

	return resp, nil
}
