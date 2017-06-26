package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Receipt struct {
}

func ReceiptIdGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get receipt with the specified ID
	var Receipt models.Declartion

	if db.Where("ID = ?", id).Find(&Receipt).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	} else {
		js, err := json.Marshal(Receipt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
	}
}

func ReceiptIdImageGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func ReceiptPost(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}
