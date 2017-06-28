package dbActions

import (
	"testing"

	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func TestCreateUser(t *testing.T) {
	dbLocation := ":memory:" //"./db/declarations.sqlite"
	db, _ := gorm.Open("sqlite3", dbLocation)
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Declaration{}, &models.Receipt{}, &models.Project{}, &models.DeclarationStatus{})

	CreateUser(&models.User{Email: "barld@barld.nl", FirstName: "Barld", LastName: "Boot", Password: "Secret"}, db)

	var users []models.User
	db.Find(&users)
	log.Print(users)
}
