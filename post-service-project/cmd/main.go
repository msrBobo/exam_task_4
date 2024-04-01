package main

import (
	//"context"
	"context"
	"encoding/json"
	"exam_task_4/post-service-project/config"
	pb "exam_task_4/post-service-project/genproto/post-service"
	"fmt"
	"net"

	c "exam_task_4/post-service-project/queue/rabbitmq/consumer"

	_ "github.com/lib/pq"

	// "post-service/pkg/db"
	"exam_task_4/post-service-project/pkg/db"
	"exam_task_4/post-service-project/pkg/logger"
	service "exam_task_4/post-service-project/service"
	"exam_task_4/post-service-project/service/grpc_client"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post-service")
	defer logger.Cleanup(log)

	createPostConsumer, err := c.NewRabbitMQConsumer("amqp://guest:guest@rabbitmq:5672/", "createPost")
	if err != nil {
		fmt.Printf("Error while initializing createUser consumer: %v", err)
		return
	}
	defer createPostConsumer.Close()
	//****************<<CONNECT FOR MONGOSH>>**********************//
	// log.Info("main mongoConfig",
	// 	logger.String("mongoURI", cfg.MongoURI),
	// 	logger.String("database", cfg.MongoDatabase))

	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	// if err != nil {
	// 	log.Fatal("MongoDB connection error", logger.Error(err))
	// }
	// defer func() {
	// 	if err := client.Disconnect(context.Background()); err != nil {
	// 		log.Error("Error disconnecting from MongoDB", logger.Error(err))
	// 	}
	// }()

	//****************<<CONNECT FOR POSTGRES>>**********************//
	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("Error creating gRPC client", logger.Error(err))
	}

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
		var post pb.Post
		err := json.Unmarshal(message, &post)
		if err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			return
		}

		postServiceP := service.NewUserServicePostgres(connPostgres, log, grpcClient)
		postServiceP.CreatePost(context.Background(), &post)

		// postServiceM := service.NewPostService(client, log, grpcClient)
		// postServiceM.CreatePost(context.Background(), &post)
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
	pb.RegisterPostServiceServer(postgres, userServicePostgres)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := postgres.Serve(RPCPort); err != nil {
		log.Fatal("Error while serving: %v", logger.Error(err))
	}

	postServiceMongo := service.NewUserServicePostgres(connPostgres, log, grpcClient)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postServiceMongo)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while serving: %v", logger.Error(err))
	}

}
