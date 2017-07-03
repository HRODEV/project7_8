package project7_8

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
)

func DeclarationsGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	var declarations []models.Declaration
	dbActions.GetDeclarationsForUser(utils.currentUser.ID, &declarations, utils.db)

	return &declarations
}

func DeclarationsIdDelete(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	dbActions.DeleteDeclarationById(uint(id), utils.db)
	w.WriteHeader(http.StatusNoContent)

	return nil
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
	var declaration models.Declaration
	dbActions.GetDeclarationById(uint(id), &declaration, utils.db)

	if declaration.ID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return nil
	} else if declaration.ID != utils.currentUser.ID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return nil
	}

	return &declaration
}

func DeclarationsIdGetImage(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Get declaration with the specified ID
	var declaration models.Declaration
	dbActions.GetDeclarationById(uint(id), &declaration, utils.db)

	if declaration.ID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return nil
	} else if declaration.ID != utils.currentUser.ID {
		http.Error(w, "Not found", http.StatusForbidden)
		return nil
	}

	img, err := os.Open(declaration.Receipt.ImagePath)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)

	return nil
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
	dbError := dbActions.UpdateDeclarationById(uint(id), &declaration, utils.db)

	if dbError != nil {
		if dbError.Error() == "not found" {
			http.Error(w, dbError.Error(), http.StatusNotFound)
			return nil
		}

		http.Error(w, dbError.Error(), http.StatusInternalServerError)
		return nil
	}

	// Get the modified declaration to return
	dbActions.GetDeclarationById(uint(id), &declaration, utils.db)

	return &declaration
}

func DeclarationsPost(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Convert request body to interface
	var declaration models.Declaration
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Append userid
	declaration.UserID = utils.currentUser.ID

	// Create declaration
	dbActions.CreateDeclaration(&declaration, utils.db)

	return &declaration
}
