package models

import (
	"time"
)

type Blog struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	UserID     int    `gorm:"not null;index:idx_user"`
	Title      string `gorm:"size:100;not null"`
	Article    string `gorm:"type:text;not null"` //博客正文
	CreateTime time.Time
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

func (Blog) TableName() string {
	return "blog"
}
