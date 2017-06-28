package models

type Project struct {
	ID   int    `gorm:"column:ID"`
	Name string `gorm:"column:Name"`
}

func (Project) TableName() string {
	return "Project"
}
