package storage

import (
	mongodb "exam_task_4/post-service-project/storage/mongosh"
	"exam_task_4/post-service-project/storage/postgres"
	"exam_task_4/post-service-project/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Post() repo.PostStorageI
}

type storageMongo struct {
	client   *mongo.Client
	postRepo repo.PostStorageI
}

func (s storageMongo) Post() repo.PostStorageI {
	return s.postRepo
}

func NewStorageMongo(client *mongo.Client) *storageMongo {
	collection := client.Database("postdb").Collection("posts")
	return &storageMongo{
		client:   client,
		postRepo: mongodb.NewPostRepo(collection),
	}
}

type storagePostgres struct {
	db               *sqlx.DB
	postRepoPostgres repo.PostStorageI
}

func (s storagePostgres) Post() repo.PostStorageI {
	return s.postRepoPostgres
}

func NewStoragePostgres(db *sqlx.DB) *storagePostgres {
	return &storagePostgres{db, postgres.NewPostRepoPostgres(db)}
}
