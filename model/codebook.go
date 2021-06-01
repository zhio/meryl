package model

import (
	"gorm.io/gorm"
	"meryl/util"
	"os"
)

type CodeBook struct {
	gorm.Model
	Title    string
	Alias    string
	Username string
	Code     string
	Notes    string `gorm:"size:1000"`
	Status   string
}

// EncryptCode 加密密码
func (codebook *CodeBook) EncryptCode(code string) error {
	AesKey := []byte(os.Getenv("KEY_SECRET")) // 对称秘钥长度必须是16的倍数
	encrypted, err := util.AesEncrypt([]byte(code), AesKey)
	if err != nil {
		return err
	}
	//codebook.Code = base64.StdEncoding.EncodeToString(encrypted)
	codebook.Code = string(encrypted)
	return nil
}

// DecryptCode 解密密码
func (codebook *CodeBook) DecryptCode() string {
	AesKey := []byte(os.Getenv("KEY_SECRET")) // 对称秘钥长度必须是16的倍数
	origin, err := util.AesDecrypt([]byte(codebook.Code), AesKey)
	if err != nil {
		util.Log().Error("解码错误: %v", err)
	}
	return string(origin)
}
