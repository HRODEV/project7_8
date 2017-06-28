package dbActions

import (
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUserByID(id int, db *gorm.DB) (user models.User) {
	db.Find(&user)
	return
}

func CreateUser(user *models.User, db *gorm.DB) {
	//db.NewRecord(*user)
	db.Save(user)
}
