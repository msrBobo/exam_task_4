package storage

import (
	mongodb "exam_task_4/comment-service-project/storage/mongosh"
	"exam_task_4/comment-service-project/storage/postgres"
	"exam_task_4/comment-service-project/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storageMongo struct {
	client      *mongo.Client
	commentRepo repo.CommentStorageI
}

func (s storageMongo) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func NewStorageMongo(client *mongo.Client) *storageMongo {
	collection := client.Database("commentdb").Collection("comments")
	return &storageMongo{
		client:      client,
		commentRepo: mongodb.NewCommentRepo(collection),
	}
}

type storagePostgres struct {
	db                  *sqlx.DB
	commentRepoPostgres repo.CommentStorageI
}

func (s storagePostgres) Comment() repo.CommentStorageI {
	return s.commentRepoPostgres
}

func NewStoragePostgres(db *sqlx.DB) *storagePostgres {
	return &storagePostgres{db, postgres.NewCommentRepoPostgres(db)}
}
