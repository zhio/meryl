package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"meryl/serializer"
	"meryl/service"
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

func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
