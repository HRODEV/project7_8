package dbActions

import (
	"errors"
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUserByID(id uint, user *models.User, db *gorm.DB) {
	db.First(user, id)
}

func CreateUser(user *models.User, db *gorm.DB) error {
	if user.IsValid() {
		db.Create(user)
		return nil
	} else {
		return errors.New("No valid password or email given. Pasword must be at least 8 characters")
	}

}
