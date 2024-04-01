package storage

import (
	mongodb "exam_task_4/user-service-project/storage/mongosh"
	"exam_task_4/user-service-project/storage/postgres"
	"exam_task_4/user-service-project/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storageMongo struct {
	client        *mongo.Client
	userRepoMongo repo.UserStorageI
}

func (s storageMongo) User() repo.UserStorageI {
	return s.userRepoMongo
}

func NewStorageMongo(client *mongo.Client) *storageMongo {
	collection := client.Database("userdb").Collection("users")
	return &storageMongo{
		client:        client,
		userRepoMongo: mongodb.NewUserRepoMongo(collection),
	}
}

type storagePostgres struct {
	db               *sqlx.DB
	userRepoPostgres repo.UserStorageI
}

func (s storagePostgres) User() repo.UserStorageI {
	return s.userRepoPostgres
}

func NewStoragePostgres(db *sqlx.DB) *storagePostgres {
	return &storagePostgres{db, postgres.NewUserRepoPostgres(db)}
}
