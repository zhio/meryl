package service

import (
	"math/rand"
	"meryl/serializer"
	"strconv"
	"time"
)

// 密码生成服务
type MakeNewPassWordService struct {
	Length  string `json:"length"`
	NumStr  string `json:"num_str"`
	CharStr string `json:"char_str"`
	SpecStr string `json:"spec_str"`
}

const (
	NumStr  = "0123456789"
	CharStr = "abcdefglijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecStr = "!@#$%^&*()[]_+-=,."
)

// Create 创建密码
func (service *MakeNewPassWordService) Create() serializer.Response {
	length, err := strconv.Atoi(service.Length)
	if err != nil {
		return serializer.ParamErr("请输入正确的参数", err)
	}
	var passwd []byte = make([]byte, length, length)
	var sourceStr string = ""
	if service.NumStr != "" {
		sourceStr += NumStr
	}
	if service.CharStr != "" {
		sourceStr += CharStr
	}
	if service.SpecStr != "" {
		sourceStr += SpecStr
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}

	return serializer.Response{
		Code: 200,
		Data: string(passwd),
	}
}
