package api

import (
	"github.com/gin-gonic/gin"
	service2 "meryl/service"

	"net/http"
)

func MakeNewKey(c *gin.Context) {
	var service service2.MakeNewPassWordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
