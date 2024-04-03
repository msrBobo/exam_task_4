package config

import (
	"os"
	"github.com/spf13/cast"
)

type Config struct {
	Environment         string
	MongoURI            string
	MongoDatabase       string
	PostgresHost		string
	PostgresPort		string
	PostgresUser		string
	PostgresPassword	string
	PostgresDatabase	string
	LogLevel            string
	RPChost 			string
	RPCPort				string
	PostServiceHost    	string
	PostServicePort    	string
	CommentServiceHost 	string
	CommentServicePort 	string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.MongoURI = cast.ToString(getOrReturnDefault("MONGO_URI", "mongodb://mongodb:27018"))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "userdb"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "postgresdb"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "userdb"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post-service-project"))
	c.PostServicePort = cast.ToString(getOrReturnDefault("POST_SERVICE_PORT", "3030"))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment-service-project"))
	c.CommentServicePort = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_PORT", "4040"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPChost = cast.ToString(getOrReturnDefault("user-service_HOST", "user-service-project"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":2020"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
