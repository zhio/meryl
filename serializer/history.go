package serializer

import "meryl/model"

// History 历史记录序列化器
type History struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Alias    string `json:"Alias"`
	CodeID   string `json:"code_id"`
	Username string `json:"username"`
	Code     string `json:"code"`
	Notes    string `json:"Notes"`
	Status   string `json:"status"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func BuildHistory(item model.History) History {
	return History{
		ID:       item.ID,
		Title:    item.Title,
		Alias:    item.Alias,
		Username: item.Username,
		Code:     item.DecryptCode(),
		Notes:    item.Notes,
		Status:   item.Status,
		CreateAt: item.CreatedAt.Unix(),
		UpdateAt: item.UpdatedAt.Unix(),
	}
}
func BuildHistories(items []model.History) (histories []History) {
	for _, item := range items {
		history := BuildHistory(item)
		histories = append(histories, history)
	}
	return histories
}
