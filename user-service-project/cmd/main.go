package main

import (
	"context"
	"encoding/json"
	"exam_task_4/user-service-project/config"
	pbu "exam_task_4/user-service-project/genproto/user-service"
	c "exam_task_4/user-service-project/queue/rabbitmq/consumer"
	"exam_task_4/user-service-project/service"
	"fmt"
	"net"

	_ "github.com/lib/pq"

	"exam_task_4/user-service-project/pkg/db"
	"exam_task_4/user-service-project/pkg/logger"
	"exam_task_4/user-service-project/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	createUserConsumer, err := c.NewRabbitMQConsumer("amqp://guest:guest@rabbitmq:5673/", "createUser")
	if err != nil {
		fmt.Printf("Error while initializing createUser consumer: %v", err)
		return
	}
	defer createUserConsumer.Close()

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

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("Error creating gRPC client", logger.Error(err))
	}

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
		var user pbu.User
		err := json.Unmarshal(message, &user)
		if err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			return
		}

		userServiceP := service.NewUserServicePostgres(connPostgres, log, grpcClient)
		userServiceP.CreateUser(context.Background(), &user)

		// userServiceM := service.NewUserServiceMongo(client, log, grpcClient)
		// userServiceM.CreateUser(context.Background(), &user)
	}
	// Start Consuming Messages
	err = createUserConsumer.ConsumeMessage(handleMessage)
	if err != nil {
		fmt.Printf("Error while consuming createUser messages: %v", err)
	}
	
	userServicePostgres := service.NewUserServicePostgres(connPostgres, log, grpcClient)
	RPCPort, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	postgres := grpc.NewServer()
	pbu.RegisterUserServiceServer(postgres, userServicePostgres)
	log.Info("main: server running with postgres",
		logger.String("port", cfg.RPCPort))
	if err := postgres.Serve(RPCPort); err != nil {
		log.Fatal("Error while serving: %v", logger.Error(err))
	}

}
