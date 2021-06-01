package serializer

import "meryl/model"

// User 用户序列化器
type CodeBook struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Alias    string `json:"alias"`
	Username string `json:"username"`
	Code     string `json:"code"`
	Nodes    string `json:"nodes"`
	Status   string `json:"status"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func BuildCodeBook(codebook model.CodeBook) CodeBook {
	return CodeBook{
		ID:       codebook.ID,
		Title:    codebook.Title,
		Alias:    codebook.Alias,
		Username: codebook.Username,
		Code:     codebook.DecryptCode(),
		Nodes:    codebook.Notes,
		Status:   codebook.Status,
		CreateAt: codebook.CreatedAt.Unix(),
		UpdateAt: codebook.UpdatedAt.Unix(),
	}
}

func BuildCodeBookResponse(codebook model.CodeBook) Response {
	return Response{
		Data: BuildCodeBook(codebook),
	}
}
