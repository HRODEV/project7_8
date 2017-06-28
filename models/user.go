package models

type User struct {
	ID        int    `gorm:"column:ID"`
	Email     string `gorm:"column:Email"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Password  string `gorm:"column:Password"`
}

func (User) TableName() string {
	return "User"
}
