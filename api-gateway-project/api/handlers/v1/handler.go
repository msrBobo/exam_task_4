package v1

import (
	"errors"
	token "exam_task_4/api-gateway-project/api/tokens"
	"exam_task_4/api-gateway-project/config"
	"exam_task_4/api-gateway-project/pkg/logger"
	"exam_task_4/api-gateway-project/queue/rabbitmq/producer"
	"exam_task_4/api-gateway-project/service"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	jwthandler     token.JWTHandler
	log            logger.Logger
	serviceManager service.IServiceManager
	cfg            config.Config
	Writer         producer.RabbitMQProducer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Jwthandler     token.JWTHandler
	Logger         logger.Logger
	ServiceManager service.IServiceManager
	Cfg            config.Config
	Enforcer       casbin.Enforcer
	Writer         producer.RabbitMQProducer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		jwthandler:     c.Jwthandler,
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		Writer:         c.Writer,
	}
}

func ParsePageQueryParam(c *gin.Context) (int, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, fmt.Errorf("page should be a positive number")
	}
	if page == 0 {
		return 1, nil
	}

	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (int, error) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("page size should be a positive number")
	}

	if limit == 0 {
		return 10, nil
	}

	return limit, nil
}

const (
	ErrorCodeInvalidURL          = "INVALID_URL"
	ErrorCodeInvalidJSON         = "INVALID_JSON"
	ErrorCodeInvalidParams       = "INVALID_PARAMS"
	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrorCodeUnauthorized        = "UNAUTHORIZED"
	ErrorCodeAlreadyExists       = "ALREADY_EXISTS"
	ErrorCodeNotFound            = "NOT_FOUND"
	ErrorCodeInvalidCode         = "INVALID_CODE"
	ErrorBadRequest              = "BAD_REQUEST"
	ErrorInvalidCredentials      = "INVALID_CREDENTIALS"
	StatusMethodNotAllowed       = "METHOD_NOT_ALLOWED"
	ErrorValidationError         = "VALIDATION_ERROR"
)
