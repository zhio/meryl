package model

import (
	"encoding/base64"
	"gorm.io/gorm"
	"meryl/util"
	"os"
)

type History struct {
	gorm.Model
	CodeId   uint
	Title    string
	Alias    string
	Username string
	Code     string
	Notes    string `gorm:"size:1000"`
	Status   string
}

// DecryptCode 解密密码
func (history *History) DecryptCode() string {
	AesKey := []byte(os.Getenv("KEY_SECRET")) // 对称秘钥长度必须是16的倍数
	bytesPass, err := base64.StdEncoding.DecodeString(history.Code)
	origin, err := util.AesDecrypt(bytesPass, AesKey)
	if err != nil {
		util.Log().Error("解码错误: %v", err)
	}
	return string(origin)
}
