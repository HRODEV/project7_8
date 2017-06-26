package main

import (
	sw "github.com/HRODEV/project7_8/routes"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open("sqlite3", "./db/declarations")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	router := sw.NewRouter(db)

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
