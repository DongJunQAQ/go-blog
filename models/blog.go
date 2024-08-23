package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserID  int    `gorm:"not null;index:idx_user"`
	Title   string `gorm:"size:100;not null"`
	Article string `gorm:"type:text;not null"` //博客正文
}

func (Blog) TableName() string {
	return "blog"
}
