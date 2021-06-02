package codebook

import (
	"meryl/model"
	"meryl/serializer"
)

type DeleteCodeBookService struct {
}

func (service *DeleteCodeBookService) Delete(id string) serializer.Response {
	var codebook model.CodeBook
	err := model.DB.First(&codebook, id).Error
	if err != nil {
		return serializer.Err(
			serializer.CodeNotFound,
			"视频不存在",
			err,
		)
	}

	err = model.DB.Delete(&codebook).Error
	if err != nil {
		return serializer.DBErr("", err)
	}
	return serializer.Response{
		Msg: "删除成功",
	}
}
