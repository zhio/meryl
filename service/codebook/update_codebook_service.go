package codebook

import (
	"encoding/json"
	"meryl/model"
	"meryl/serializer"
)

// UpdateCodeBookService 更新密码本服务
type UpdateCodeBookService struct {
	Title       string `form:"title" json:"title" binding:"min=2,max=30"`
	Alias       string `form:"Alias" json:"Alias" binding:"max=200"`
	Username    string `form:"username" json:"username" binding:"min=2,max=30"`
	Code        string `form:"code" json:"code" binding:"required,min=8,max=40"`
	CodeConfirm string `form:"code_confirm" json:"code_confirm" binding:"min=8,max=40"`
	Notes       string `form:"Notes" json:"Notes"`
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
	model.DB.Model(&model.CodeBook{}).Where("Alias = ?", service.Alias).Count(&count)

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
	if service.Alias != "" {
		codebook.Alias = service.Alias
	}
	if service.Notes != "" {
		codebook.Notes = service.Notes
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

//包含历史版本的更新操作
func (service *UpdateCodeBookService) UpdateByHistory(id string) serializer.Response {
	var codebook model.CodeBook

	err := model.DB.First(&codebook, id).Error
	if err != nil {
		return serializer.Err(
			serializer.CodeNotFound,
			"密码本不存在",
			err,
		)
	}

	history := model.History{
		CodeId:   codebook.ID,
		Title:    codebook.Title,
		Alias:    codebook.Alias,
		Username: codebook.Username,
		Code:     codebook.Code,
		Notes:    codebook.Notes,
		Status:   model.Inactive,
	}

	if err := model.DB.Create(&history).Error; err != nil {
		return serializer.ParamErr("添加历史记录失败", err)
	}

	bytes, _ := json.Marshal(service)
	var dataModel model.CodeBook
	json.Unmarshal(bytes, &dataModel)

	// 加密密码
	if err := dataModel.EncryptCode(service.Code); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	updates := model.DB.Model(&codebook).Updates(&dataModel)

	if updates.RowsAffected > 0 {
		return serializer.Response{
			Code: 200,
			Msg:  "更新成功",
			Data: "ok",
		}
	} else {
		return serializer.Response{
			Code: 500,
			Msg:  "更新失败",
			Data: updates.Error,
		}
	}
}
