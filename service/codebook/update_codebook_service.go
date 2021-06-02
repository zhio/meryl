package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

// UpdateCodeBookService 更新密码本服务
type UpdateCodeBookService struct {
	Title       string `form:"title" json:"title" binding:"min=2,max=30"`
	ALias       string `form:"alias" json:"alias" binding:"max=200"`
	Username    string `form:"username" json:"username" binding:"min=2,max=30"`
	Code        string `form:"code" json:"code" binding:"required,min=8,max=40"`
	CodeConfirm string `form:"code_confirm" json:"code_confirm" binding:"min=8,max=40"`
	Nodes       string `form:"nodes" json:"nodes"`
}

// valid 验证表单
func (service *UpdateCodeBookService) valid() *serializer.Response {
	if service.CodeConfirm != service.Code {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := int64(0)
	model.DB.Model(&model.CodeBook{}).Where("alias = ?", service.ALias).Count(&count)

	if count > 0 {
		return &serializer.Response{
			Code: 400001,
			Msg:  "别名被占用",
		}
	}
	return nil
}

// Update 更新密码本
func (service *UpdateCodeBookService) Update(id string) serializer.Response {
	var codebook model.CodeBook
	err := model.DB.First(&codebook, id).Error
	if err != nil {
		return serializer.Err(
			serializer.CodeNotFound,
			"视频不存在",
			err,
		)
	}
	if err := service.valid(); err != nil {
		return *err
	}

	if err := codebook.EncryptCode(service.Code); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	if service.Title != "" {
		codebook.Title = service.Title
	}
	if service.ALias != "" {
		codebook.Alias = service.ALias
	}
	if service.Nodes != "" {
		codebook.Notes = service.Nodes
	}
	if service.Username != "" {
		codebook.Username = service.Username
	}

	err = model.DB.Save(&codebook).Error
	if err != nil {
		return serializer.DBErr("", err)
	}
	return serializer.BuildCodeBookResponse(codebook)
}
