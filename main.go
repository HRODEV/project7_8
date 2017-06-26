package main

import (
	sw "github.com/HRODEV/project7_8/routes"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	dbLocation := "./db/declarations"

	// First argument is de db location
	if len(args) > 1 {
		if args[1] != "" {
			dbLocation = os.Args[1]
		}
	}

	db, err := gorm.Open("sqlite3", dbLocation)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	router := sw.NewRouter(db)

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
