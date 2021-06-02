package api

import (
	"github.com/gin-gonic/gin"
	"meryl/service/codebook"
	"net/http"
)

// CreateCode 创建密码本
func CreateCode(c *gin.Context) {
	var service codebook.CreateCodeBookService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ShowCode 密码本详情
func ShowCode(c *gin.Context) {
	var service codebook.ShowCodeBookService
	var res = service.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// ListCode 密码列表
func ListCode(c *gin.Context) {
	var service codebook.ListCodeBookService
	var res = service.List()
	c.JSON(http.StatusOK, res)
}

// UpdateCode 更新密码
func UpdateCode(c *gin.Context) {
	var service codebook.UpdateCodeBookService
	if err := c.ShouldBind(&service); err == nil {
		var res = service.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// DeleteCode 删除密码
func DeleteCode(c *gin.Context) {
	var service codebook.DeleteCodeBookService
	var res = service.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
