package models

type Receipt struct {
	ID            uint `gorm:"primary_key;AUTO_INCREMENT"`
	ImagePath     string
	Data          string `gorm:"type:text"`
	Declaration   *Declaration
	DeclarationID uint `sql:"type:integer REFERENCES declarations(id)"`
}
