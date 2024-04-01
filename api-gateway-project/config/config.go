package config

import (
	"github.com/spf13/cast"
	"os"
)
type Config struct {
	Environment         string
	MockServiceHost 	string
	MockServicePort 	int
	UserServiceHost     string
	UserServicePort     int
	PostServiceHost     string
	PostServicePort     int
	CommentServiceHost  string
	CommentServicePort  int
	CtxTimeout          int
	LogLevel            string
	HTTPPort            string
	AuthConfigPatch     string
	CSVFile             string
	SignInKey           string
	AccessTokenTimeout  int
}

func Load() Config {
	c := Config{}

	c.MockServiceHost = cast.ToString(getOrReturnDefault("LINK_SERVICE_HOST", "localhost"))
	c.MockServicePort = cast.ToInt(getOrReturnDefault("LINK_SERVICE_PORT", 1010))

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":7007"))
	
	c.SignInKey = cast.ToString(getOrReturnDefault("SIGN_IN_KEY", "JNDKmklmoemkMEK223!@#0#komklmem"))
	c.AccessTokenTimeout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMEOUT", 300))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user-service-project"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 2020))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post-service-project"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 3030))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment-service-project"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 4040))

	c.AuthConfigPatch = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATCH", "auth.conf"))
	c.CSVFile = cast.ToString(getOrReturnDefault("CSV_FILE_PATCH", "auth.csv"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
