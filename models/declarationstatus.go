package models

type DeclarationStatus struct {
	ID               int    `gorm:"column:ID"`
	Status           string `gorm:"column:Status"`
	DateModified     string `gorm:"column:DateModified"`
	DeclarationID    int    `gorm:"column:DeclarationID"`
	ModifiedByUserID int    `gorm:"column:ModifiedByUserID"`
}

func (DeclarationStatus) TableName() string {
	return "DeclarationStatus"
}
