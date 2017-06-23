package project7_8

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Declarations struct {
	ID          int
	Title       string
	Description string
}

func DeclarationsGet(w http.ResponseWriter, r *http.Request) {
	declaration := Declarations{1, "Reiskosten", "Reiskosten naar de klant"}

	js, err := json.Marshal(declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func DeclarationsIdDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeclarationsIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varx, err := strconv.Atoi(vars["id"])
	declaration := Declarations{varx, "Reiskosten", "Reiskosten naar de klant"}

	js, err := json.Marshal(declaration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func DeclarationsIdPatch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeclarationsPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
