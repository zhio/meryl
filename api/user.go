package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"meryl/cache"
	"meryl/serializer"
	"meryl/service"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var servicer service.UserRegisterService
	if err := c.ShouldBind(&servicer); err == nil {
		res := servicer.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var servicer service.UserLoginService
	if err := c.ShouldBind(&servicer); err == nil {
		res := servicer.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserTokenRefresh(c *gin.Context) {
	user := CurrentUser(c)
	var service service.UserTokenRefreshService
	res := service.Refresh(c, user)
	c.JSON(http.StatusOK, res)
}
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(http.StatusOK, res)
}

func UserLogout(c *gin.Context) {
	//使用token登出方式
	token := c.GetHeader("x-token")
	if token != "" {
		_ = cache.DelUserToken(token)
	} else {
		//使用session登出方式
		s := sessions.Default(c)
		s.Clear()
		s.Save()
	}

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
