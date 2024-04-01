package middleware

import (
	jwt "exam_task_4/api-gateway-project/api/tokens"
	"exam_task_4/api-gateway-project/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Auth(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/v1/login" || ctx.Request.URL.Path == "/v1/register" || ctx.Request.URL.Path == "/v1/Verification" || ctx.Request.URL.Path == "/v1/swagger/swagger/doc.json" || ctx.Request.URL.Path == "/v1/swagger/index.html" || ctx.Request.URL.Path == "/v1/swagger/swagger-ui.css" || ctx.Request.URL.Path == "/v1/swagger/swagger-ui-bundle.js" || ctx.Request.URL.Path == "/v1/swagger/swagger-ui-standalone-preset.js" {
		ctx.Next()
		return
	}
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized request",
		})
		return
	}

	claims, err := jwt.ExtractClaim(token, []byte(config.Load().SignInKey))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
	if cast.ToString(claims["role"]) == "user" {
		if ctx.Request.URL.Path == "/v1/users" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "no access for regular users",
			})
		}

		return
	} else if cast.ToString(claims["role"]) == "admin" {
		ctx.Next()

	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized request",
	})
}
