package models

import "time"

type Declaration struct {
	ID                  uint `gorm:"primary_key;AUTO_INCREMENT"`
	Title               string
	TotalPrice          float32
	VATPrice            float32
	Date                time.Time
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
