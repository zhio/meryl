package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

// CreateCodeBookService 密码本创建服务
type CreateCodeBookService struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Alias       string `form:"Alias" json:"Alias" binding:"max=200"`
	Username    string `form:"username" json:"username" binding:"required,min=2,max=30"`
	Code        string `form:"code" json:"code" binding:"required,min=8,max=40"`
	CodeConfirm string `form:"code_confirm" json:"code_confirm" binding:"required,min=8,max=40"`
	Notes       string `form:"Notes" json:"Notes"`
}

// valid 验证表单
func (service *CreateCodeBookService) valid() *serializer.Response {
	if service.CodeConfirm != service.Code {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := int64(0)
	model.DB.Model(&model.CodeBook{}).Where("Alias = ?", service.Alias).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 400001,
			Msg:  "别名被占用",
		}
	}
	return nil
}

// Create 创建密码本
func (service *CreateCodeBookService) Create() serializer.Response {
	codebook := model.CodeBook{
		Title:    service.Title,
		Alias:    service.Alias,
		Username: service.Username,
		Notes:    service.Notes,
		Status:   model.Active,
	}
	if err := service.valid(); err != nil {
		return *err
	}
	// 加密密码
	if err := codebook.EncryptCode(service.Code); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	// 创建密码本
	if err := model.DB.Create(&codebook).Error; err != nil {
		return serializer.ParamErr("添加密码本失败", err)
	}

	return serializer.BuildCodeBookResponse(codebook)
}
