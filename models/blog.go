package models

import (
	"time"
)

type Blog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"not null;index:idx_user"`
	Title      string    `gorm:"size:100;not null"`
	Article    string    `gorm:"type:text;not null"` //博客正文
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

func (Blog) TableName() string {
	return "blog"
}
