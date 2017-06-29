package dbActions

import (
	"errors"
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
)

func GetDeclarations(declaration *[]models.Declaration, db *gorm.DB) {
	db.Find(&declaration)
}

func GetDeclarationById(id uint, declaration *models.Declaration, db *gorm.DB) {
	db.First(declaration, id)
}

func CreateDeclaration(declaration *models.Declaration, db *gorm.DB) error {
	if declaration.IsValid() {
		db.Create(declaration)
		return nil
	} else {
		return errors.New("Declaration Struct not valid")
	}
}

func UpdateDeclarationById(id uint, declaration *models.Declaration, db *gorm.DB) error {
	if declaration.IsValid() {
		currentDeclarations := models.Declaration{}
		db.First(&currentDeclarations, id)
		db.Model(&currentDeclarations).Update(&declaration)
		return nil
	} else {
		return errors.New("Declaration Struct not valid")
	}
}
