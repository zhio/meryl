package serializer

import "meryl/model"

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
