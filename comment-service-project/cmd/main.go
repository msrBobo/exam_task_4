package main

import (
	"context"
	"encoding/json"
	"exam_task_4/comment-service-project/config"
	pb "exam_task_4/comment-service-project/genproto/comment-service"
	c "exam_task_4/comment-service-project/queue/rabbitmq/consumer"
	"fmt"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"context"
	"exam_task_4/comment-service-project/pkg/db"
	"exam_task_4/comment-service-project/pkg/logger"
	service "exam_task_4/comment-service-project/service"
	"exam_task_4/comment-service-project/service/grpc_client"
	"net"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("Error creating gRPC client", logger.Error(err))
	}
	createPostConsumer, err := c.NewRabbitMQConsumer("amqp://guest:guest@rabbitmq:5672/", "createComment")
	if err != nil {
		fmt.Printf("Error while initializing createUser consumer: %v", err)
		return
	}
	defer createPostConsumer.Close()
	//****************<<CONNECT FOR MONGOSH>>**********************//
	log.Info("main mongoConfig",
		logger.String("mongoURI", cfg.MongoURI),
		logger.String("database", cfg.MongoDatabase))

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error", logger.Error(err))
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Error("Error disconnecting from MongoDB", logger.Error(err))
		}
	}()

	//****************<<CONNECT FOR POSTGRES>>**********************//
	log.Info("main postgresConfig",
		logger.String("postgresHost", cfg.PostgresHost),
		logger.String("postgresPort", cfg.PostgresPort),
		logger.String("postgresDatabase", cfg.PostgresDatabase),
	)

	connPostgres, err := db.ConnectToPostgresDB(cfg)
	if err != nil {
		log.Fatal("Postgres connection error", logger.Error(err))
	}

	handleMessage := func(message []byte) {
		var comment pb.Comment
		err := json.Unmarshal(message, &comment)
		if err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			return
		}

		commentServiceP := service.NewUserServicePostgres(connPostgres, log, grpcClient)
		commentServiceP.CreateComment(context.Background(), &comment)

		commentServiceM := service.NewCommentServiceMongo(client, log, grpcClient)
		commentServiceM.CreateComment(context.Background(), &comment)
	}
	// Start Consuming Messages
	err = createPostConsumer.ConsumeMessage(handleMessage)
	if err != nil {
		fmt.Printf("Error while consuming createUser messages: %v", err)
	}

	userServicePostgres := service.NewUserServicePostgres(connPostgres, log, grpcClient)
	RPCPort, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	postgres := grpc.NewServer()
	pb.RegisterCommentServiceServer(postgres, userServicePostgres)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := postgres.Serve(RPCPort); err != nil {
		log.Fatal("Error while serving: %v", logger.Error(err))
	}

	commentServiceMongo := service.NewUserServicePostgres(connPostgres, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, commentServiceMongo)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while serving: %v", logger.Error(err))
	}

}
