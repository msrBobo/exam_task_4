package main

import (
	"exam_task_4/api-gateway-project/api"
	"exam_task_4/api-gateway-project/config"
	"exam_task_4/api-gateway-project/pkg/logger"
	"exam_task_4/api-gateway-project/queue/rabbitmq/producer"
	"exam_task_4/api-gateway-project/service"
	"fmt"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	serviceManager, err := service.NewServiceManager(&cfg)
	if err != nil {
		log.Fatal("gRPC dial error", logger.Error(err))
	}

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPatch, cfg.CSVFile)
	if err != nil {
		log.Fatal("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()

	if err != nil {
		log.Fatal("casbin error load policy", logger.Error(err))
		return
	}
	
	writer, err := producer.NewRabbitMQProducer("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		fmt.Println(err)
		log.Fatal("error creating RabbitMQ producer")
	}
	
	defer writer.Close()

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	server := api.New(&api.Option{
		Conf:           cfg,
		Logger:         log,
		CasbinEnforcer: casbinEnforcer,
		ServiceManager: serviceManager,
		Writer:         *writer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		return
	}
}
