package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

// ListCodeBookService 密码本列表服务
type ListCodeBookService struct {
}

// List 展示全部密码本
func (service *ListCodeBookService) List() serializer.Response {
	var codebooks []model.CodeBook
	err := model.DB.Find(&codebooks).Error
	if err != nil {
		return serializer.DBErr("", err)
	}
	return serializer.BuildCodeBooksResponse(codebooks)
}
