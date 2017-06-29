package models

type Receipt struct {
	ID        int    `gorm:"column:ID;primary_key"`
	ImagePath string `gorm:"column:ImagePath"`
	Data      string `gorm:"column:Data"`
}

func (Receipt) TableName() string {
	return "Receipt"
}
