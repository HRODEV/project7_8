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
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func UserAuthGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not implemented yet"))
}

func UserGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	var user *models.User
	dbActions.GetUserByID(uint(id), user, utils.db)

	return user
}

//
func UserPost(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Convert request body to interface
	var user *models.User
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	dbActions.CreateUser(user, utils.db)
	return user
}

func UserProjectsGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	return errors.New("not yet implemented")
}
