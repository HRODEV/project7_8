package models

type Receipt struct {
	ID        int    `gorm:"column:ID"`
	ImagePath string `gorm:"column:ImagePath"`
	Date      string `gorm:"column:Date"`
}

func (Receipt) TableName() string {
	return "Date"
}
