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

func insertTestData(db *gorm.DB) {
	users := []models.User{
		{FirstName: "Barld", LastName: "Boot", Email: "boot@barld.nl", Password: "secrect123"},
		{FirstName: "Niels", LastName: "Van der Veer", Email: "niels@niels.nl", Password: "secrect123"},
		{FirstName: "Thom", LastName: "Overhand", Email: "thom@dodge.beer", Password: "secrect123"},
	}
	for _, user := range users {
		dbActions.CreateUser(&user, db)
	}

	declarations := []models.Declaration{
		{Title: "first declaration", TotalPrice: 95.32, VATPrice: (95.32 / 121 * 21), Date: "07-08-2016", UserID: uint(1)},
	}
	for _, declaration := range declarations {
		dbActions.CreateDeclaration(&declaration, db)
	}
}

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

	if len(args) > 2 {
		if args[2] == "--testdata" {
			insertTestData(db)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	router := sw.NewRouter(db)

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
