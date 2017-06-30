package dbActions

import (
	"testing"

	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var initialized bool = false

func init() {
	if !initialized {
		dbLocation := ":memory:" //"./db/declarations.sqlite"
		db, _ = gorm.Open("sqlite3", dbLocation)
		//db.LogMode(true)
		db.AutoMigrate(&models.User{}, &models.Declaration{}, &models.Receipt{}, &models.Project{}, &models.DeclarationStatus{})
	}
	initialized = true
}

func TestCreateUser(t *testing.T) {
	newUser := models.User{Email: "barld@barld.nl", FirstName: "Barld", LastName: "Boot", Password: "Secret12345"}
	CreateUser(&newUser, db)

	var lastUser models.User
	db.Last(&lastUser)
	if lastUser != newUser {
		t.Errorf("the last user should be %+v but was %+v", newUser, lastUser)
	}
}

func TestGetUserByID(t *testing.T) {
	newUser := models.User{Email: "barld2@barld.nl", FirstName: "Barld", LastName: "Boot", Password: "Secret12345"}
	CreateUser(&newUser, db)

	var UserByID models.User
	GetUserByID(newUser.ID, &UserByID, db)
	if UserByID != newUser {
		t.Errorf("the last user should be %+v but was %+v", newUser, UserByID)
	}
}
