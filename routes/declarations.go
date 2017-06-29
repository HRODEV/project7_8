package project7_8

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
)

func DeclarationsGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	declarations := []models.Declaration{}
	dbActions.GetDeclarations(&declarations, utils.db)

	return &declarations
}

func DeclarationsIdDelete(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	return errors.New("Not Implemented")
}

func DeclarationsIdGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Get declaration with the specified ID
	declaration := models.Declaration{}
	dbActions.GetDeclarationById(uint(id), &declaration, utils.db)

	if declaration.ID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil
	}

	return &declaration
}

func DeclarationsIdPatch(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	// Convert request body to interface
	declaration := models.Declaration{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Update declaration
	dbActions.UpdateDeclarationById(uint(id), &declaration, utils.db)

	return &declaration
}

func DeclarationsPost(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Convert request body to interface
	declaration := models.Declaration{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Create declaration
	dbActions.CreateDeclaration(&declaration, utils.db)

	return &declaration
}
