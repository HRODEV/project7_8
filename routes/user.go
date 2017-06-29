package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

func UserAuthGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func UserGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	var user models.User
	dbActions.GetUserByID(uint(id), &user, db)

	js, err := json.Marshal(user)

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
