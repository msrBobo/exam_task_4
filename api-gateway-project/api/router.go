package api

import (
	_ "exam_task_4/api-gateway-project/api/docs"
	v1 "exam_task_4/api-gateway-project/api/handlers/v1"
	token "exam_task_4/api-gateway-project/api/tokens"
	"exam_task_4/api-gateway-project/config"
	"exam_task_4/api-gateway-project/pkg/logger"
	"exam_task_4/api-gateway-project/queue/rabbitmq/producer"
	"exam_task_4/api-gateway-project/service"
	"github.com/gin-contrib/cors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Enforcer       casbin.Enforcer
	CasbinEnforcer *casbin.Enforcer
	ServiceManager service.IServiceManager
	Writer         producer.RabbitMQProducer
}

// Constructor
// @Title EXAM_TASK_4_APIS
// @version 1.0
// @description api-gateway
// @securityDefinitions.apikey BearerAuth
// @host 18.133.228.143:7007
// @in header
// @name Authorization
func New(option *Option) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	jwtHandler := token.JWTHandler{
		SignKey: option.Conf.SignInKey,
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Jwthandler:     jwtHandler,
		Writer:         option.Writer,
	})

	api := router.Group("/v1")

	api.POST("/user/create", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.GetUser)
	api.GET("/getall", handlerV1.ListUsers)
	api.GET("/users/posts", handlerV1.GetPostsByUserId)
	api.PUT("/user/update", handlerV1.UpdateUser)
	api.DELETE("/user/delete", handlerV1.DeleteUser)

	api.POST("/post/create", handlerV1.CreatePost)
	api.GET("/post/:id", handlerV1.GetPost)
	api.GET("/posts", handlerV1.GetAllPosts)
	api.GET("/postsbyuserid/:id", handlerV1.GetPostsByUserId)
	api.GET("/postsbyuseridwithcomments/:id", handlerV1.GetPostsAndCommentsByUserId)
	api.PUT("/post/update", handlerV1.UpdatePost)
	api.DELETE("/post/delete", handlerV1.DeletePost)

	api.POST("/comment/create", handlerV1.CreateComment)
	api.GET("/comment/:id", handlerV1.GetComment)
	api.GET("/comments", handlerV1.GetAllComments)
	api.GET("/commentsby/:id/", handlerV1.GetCommentsByPostId)
	api.GET("/commentsbyuser/:id/", handlerV1.GetCommentsByUserId)
	api.GET("/allpostswithcommentsandowners/", handlerV1.GetAllPostsWithCommentsAndOwners)
	api.PUT("/comment/update", handlerV1.UpdateComment)
	api.DELETE("/comment/delete", handlerV1.DeleteComment)

	url := ginSwagger.URL("swagger/doc.json")

	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
