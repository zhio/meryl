package server

import (
	"github.com/gin-gonic/gin"
	"meryl/api"
	"meryl/middleware"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//中间件 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		//用户注册
		v1.POST("user/register", api.UserRegister)

		//用户登陆
		v1.POST("user/login", api.UserLogin)

		// 需要登陆保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			//auth.GET("codebook/:id",api.ShowCode)
			auth.POST("codebooks", api.CreateCode)
			//auth.GET("codebooks",api.ListCode)
			//auth.PUT("codebook/:id",api.UpdateCode)
			//auth.DELETE("codebook/:id",api.DeleteCode)
		}
	}
	return r

}
