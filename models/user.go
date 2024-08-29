package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"` //自动编号从1开始，``内的内容被称为Tag
	Name     string `gorm:"unique;size:20"`
	Password string `gorm:"size:32"`
}

func (User) TableName() string { //设置自动创建的表名
	return "user"
}
