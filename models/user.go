package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"unique;size:20"`
	Password string `gorm:"size:32"`
}
