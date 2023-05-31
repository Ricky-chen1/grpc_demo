package middleware

import (
	"grpc_demo/app/api/handler"
	"grpc_demo/pkg/errno"
	"grpc_demo/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JwT middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			handler.BuildFailResponse(c, errno.NoToken, errno.CodeTag[errno.NoToken])
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			handler.BuildFailResponse(c, errno.TokenParseFail, errno.CodeTag[errno.TokenParseFail])
			c.Abort()
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			handler.BuildFailResponse(c, errno.TokenExpired, errno.CodeTag[errno.TokenExpired])
			c.Abort()
			return
		}

		c.Next()
	}
}
