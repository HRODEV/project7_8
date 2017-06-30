package models

import "time"

type DeclarationStatus struct {
	ID            uint `gorm:"primary_key;AUTO_INCREMENT"`
	Status        string
	DateModified  time.Time
	Declaration   *Declaration
	DeclarationID uint `sql:"type:integer REFERENCES declarations(id)"`
	User          *User
	UserId        uint `sql:"type:integer REFERENCES users(id)"`
}
