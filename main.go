package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	sw "github.com/HRODEV/project7_8/routes"
	"github.com/jinzhu/gorm"
)

func main() {
	args := os.Args
	dbLocation := "./db/declarations.sqlite"

	// First argument is de db location
	if len(args) > 1 {
		if args[1] != "" {
			dbLocation = os.Args[1]
		}
	}

	db, err := gorm.Open("sqlite3", dbLocation)
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Declaration{}, &models.Receipt{}, &models.Project{}, &models.DeclarationStatus{})

	dbActions.CreateUser(&models.User{Email: "barld@barld.nl", FirstName: "Barld", LastName: "Boot", Password: "Secret"}, db)
	//db.Model(&models.Declaration{}).
	//	AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	if err != nil {
		log.Fatal(err)
	}

	router := sw.NewRouter(db)

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
