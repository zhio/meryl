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

func BuildCodeBook(item model.CodeBook) CodeBook {
	return CodeBook{
		ID:       item.ID,
		Title:    item.Title,
		Alias:    item.Alias,
		Username: item.Username,
		Code:     item.DecryptCode(),
		Nodes:    item.Notes,
		Status:   item.Status,
		CreateAt: item.CreatedAt.Unix(),
		UpdateAt: item.UpdatedAt.Unix(),
	}
}
func BuildCodeBooks(items []model.CodeBook) (codebooks []CodeBook) {
	for _, item := range items {
		codebook := BuildCodeBook(item)
		codebooks = append(codebooks, codebook)
	}
	return codebooks
}
func BuildCodeBookResponse(codebook model.CodeBook) Response {
	return Response{
		Data: BuildCodeBook(codebook),
	}
}

func BuildCodeBooksResponse(codebooks []model.CodeBook) Response {
	return Response{
		Data: BuildCodeBooks(codebooks),
	}
}
