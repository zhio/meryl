package middleware

import (
	"github.com/gin-gonic/gin"
	"meryl/serializer"
	"meryl/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "" {
			claims, err := util.ParseToken(token)
			if err == nil && time.Now().Unix() < claims.ExpiresAt {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
