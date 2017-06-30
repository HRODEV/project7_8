package models

type Declaration struct {
	ID                  uint `gorm:"primary_key;AUTO_INCREMENT"`
	Title               string
	TotalPrice          float32
	VATPrice            float32
	Date                string
	Description         string
	Project             *Project
	ProjectID           uint `sql:"type:integer REFERENCES projects(id)"`
	StoreName           string
	Receipt             *Receipt
	ReceiptID           uint `sql:"type:integer REFERENCES receipts(id)"`
	User                *User
	UserID              uint `sql:"type:integer REFERENCES users(id)"`
	DeclarationStatusus *[]DeclarationStatus
}

func (declaration *Declaration) IsValid() bool {
	return true
}
