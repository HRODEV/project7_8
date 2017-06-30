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
	// Check if the given struct is valid
	if !declaration.IsValid() {
		return errors.New("The given declaration struct is not valid")
	}

	// Get the latest declaration to compare
	var currentDeclarations models.Declaration

	if db.First(&currentDeclarations, id).RecordNotFound() {
		return errors.New("not found")
	}

	db.Model(&currentDeclarations).Update(&declaration)

	return nil
}

func DeleteDeclarationById(id uint, db *gorm.DB) {
	db.Delete(&models.Declaration{ID: id})
}
