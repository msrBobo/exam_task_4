package db

import (
	"context"
	"exam_task_4/user-service-project/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(cfg config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectToPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}
	return connDB, nil
}
