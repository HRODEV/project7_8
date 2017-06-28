package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func UserAuthGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func UserGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var users []models.User
	db.Find(&users)

	js, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func UserPost(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func UserProjectsGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}
