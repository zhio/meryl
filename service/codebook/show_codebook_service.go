package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

// ShowCodeBookService 密码本详情服务
type ShowCodeBookService struct {
}

// Show 展示密码本
func (service *ShowCodeBookService) Show(id string) serializer.Response {
	var codebook model.CodeBook
	err := model.DB.First(&codebook, id).Error
	if err != nil {
		return serializer.DBErr(
			"密码本不存在",
			err)
	}
	return serializer.BuildCodeBookResponse(codebook)
}
