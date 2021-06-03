package model

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	CodeId   string
	Title    string
	Alias    string
	Username string
	Code     string
	Notes    string `gorm:"size:1000"`
	Status   string
}
