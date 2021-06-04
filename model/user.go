package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"meryl/cache"
	"meryl/util"
	"strconv"
	"time"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活
	Active string = "active"
	// Inactive 未激活
	Inactive string = "inactive"
	// Suspend 被封禁
	Suspend string = "suspend"
)

// GetUser 用id获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// UserID 返回string版的uid
func (user *User) UserID() string {
	return strconv.Itoa(int(user.ID))
}

// MakeToken 生成token
func (user *User) MakeToken() (string, int64, error) {
	token := util.RandStringRunes(15)
	exp := 14 * 24 * time.Hour
	tokenExpire := time.Now().Add(exp).Unix()
	if err := cache.SaveUserToken(token, user.UserID(), exp); err != nil {
		return "", 0, err
	}
	return token, tokenExpire, nil
}
