package middleware

import (
	"gitee.com/lyhuilin/pkg/errno"
	"gitee.com/lyhuilin/pkg/handler"
	"gitee.com/lyhuilin/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
