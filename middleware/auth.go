package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"meryl/cache"
	"meryl/model"
	"meryl/serializer"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid string
		token := c.GetHeader("x-token")
		if token != "" {
			uid, _ = cache.GetUserByToken(token)
		} else {
			session := sessions.Default(c)
			uid, _ = session.Get("user_id").(string)
		}

		if uid != "" {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
