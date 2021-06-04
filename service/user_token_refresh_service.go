package service

import (
	"github.com/gin-gonic/gin"
	"meryl/model"
	"meryl/serializer"
)

type UserTokenRefreshService struct {
}

func (service *UserTokenRefreshService) Refresh(c *gin.Context, user *model.User) serializer.Response {
	token, tokenExpire, err := user.MakeToken()
	if err != nil {
		return serializer.DBErr("redis错误", err)
	}
	data := serializer.BuildUser(*user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Data: data,
	}
}
