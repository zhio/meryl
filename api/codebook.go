package api

import (
	"github.com/gin-gonic/gin"
	"meryl/service/codebook"
)

// CreateCode 创建密码本
func CreateCode(c *gin.Context) {
	var service codebook.CreateCodeBookService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowCode 密码本详情
func ShowCode() {

}

// ListCode 密码列表
func ListCode() {

}

// UpdateCode 更新密码
func UpdateCode() {

}

// DeleteCode 删除密码
func DeleteCode() {

}
