package models

type Declartion struct {
	ID          int64   `gorm:"column:ID"`
	Title       string  `gorm:"column:Title"`
	TotalPrice  float32 `gorm:"column:TotalPrice"`
	VATPrice    float32 `gorm:"column:VATPrice"`
	Date        string  `gorm:"column:Date"`
	Description string  `gorm:"column:Description"`
	ProjectID   int     `gorm:"column:ProjectID"`
	StoreName   string  `gorm:"column:StoreName"`
	ReceiptID   int     `gorm:"column:ReceiptID"`
	UserID      int     `gorm:"column:UserID"`
}

func (Declartion) TableName() string {
	return "Declaration"
}
