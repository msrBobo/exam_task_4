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
	RPChost				string
	RPCPort             string
	UserServiceHost    	string
	UserServicePort    	string
	PostServiceHost 	string
	PostServicePort 	string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.MongoURI = cast.ToString(getOrReturnDefault("MONGO_URI", "mongodb://mongodb:27018"))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "commentdb"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "postgresdb"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "commentdb"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "post-service-project"))
	c.PostServicePort = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_PORT", "3030"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user-service-project"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SERVICE_PORT", "2020"))
	
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPChost = cast.ToString(getOrReturnDefault("comment-service_HOST", "comment-service-project"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":4040"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
