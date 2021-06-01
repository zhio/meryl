package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

// CreateCodeBookService 密码本创建服务
type CreateCodeBookService struct {
	Title    string `form:"title" json:"title" binding:"required,min=2,max=30"`
	ALias    string `form:"alias" json:"alias" binding:"max=200"`
	Username string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Code     string `form:"code" json:"code" binding:"required,min=8,max=40"`
	Nodes    string `form:"nodes" json:"nodes"`
}

// Create 创建密码本
func (service *CreateCodeBookService) Create() serializer.Response {
	codebook := model.CodeBook{
		Title:    service.Title,
		Alias:    service.ALias,
		Username: service.Username,
		Notes:    service.Nodes,
		Status:   model.Active,
	}

	// todo 表单校验
	// 加密密码
	if err := codebook.EncryptCode(service.Code); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	// 创建用户
	if err := model.DB.Create(&codebook).Error; err != nil {
		return serializer.ParamErr("添加密码本失败", err)
	}

	return serializer.BuildCodeBookResponse(codebook)
}
