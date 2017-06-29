package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DeclarationsGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var declarations []models.Declartion
	db.Find(&declarations)

	js, err := json.Marshal(declarations)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func DeclarationsIdDelete(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func DeclarationsIdGet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get declaration with the specified ID
	var declaration models.Declartion

	if db.Where("ID = ?", id).Find(&declaration).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	} else {
		js, err := json.Marshal(declaration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
	}
}

func DeclarationsIdPatch(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Convert request body to interface
	declaration := models.Declartion{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add primary_key to the struct
	declaration.ID = id

	// Get the current values and insert the difference
	currentDeclarations := models.Declartion{}
	db.Where("ID = ?", id).First(&currentDeclarations)
	db.Model(&currentDeclarations).Update(&declaration)

	// Render inserted object
	enc := json.NewEncoder(w)
	enc.Encode(&currentDeclarations)
}

func DeclarationsPost(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Convert request body to interface
	declaration := models.Declartion{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert into database
	db.Create(&declaration)

	// Render inserted object
	enc := json.NewEncoder(w)
	enc.Encode(&declaration)
}
