package models

type UserProject struct {
	ProjectID int `gorm:"column:ProjectID"`
	UserID    int `gorm:"column:UserID"`
}

func (UserProject) TableName() string {
	return "UserProject"
}
