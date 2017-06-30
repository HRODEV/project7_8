package models

type Project struct {
	ID           uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string
	Users        *[]User
	Declarations *[]Declaration
}
