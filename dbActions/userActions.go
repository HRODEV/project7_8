package dbActions

import (
	"errors"
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUsers(users *[]models.User, db *gorm.DB) {
	db.Find(users)
}

func GetUserByID(id uint, user *models.User, db *gorm.DB) {
	db.First(user, id)
}

func CreateUser(user *models.User, db *gorm.DB) error {
	valid, err := user.IsValid()

	if valid {
		if dbErrors := db.Create(user).GetErrors(); len(dbErrors) != 0 {
			if dbErrors[0].Error() == "UNIQUE constraint failed: users.email" {
				return errors.New("The email address you try to register is already in use")
			} else {
				return errors.New("Database error...")
			}
		}

		return nil
	} else {
		return err
	}
}
