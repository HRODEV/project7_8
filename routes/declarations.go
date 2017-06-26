package project7_8

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Declaration struct {
	ID          int
	Title       string
	Description string
}

func DeclarationsGet(w http.ResponseWriter, r *http.Request) {
	declarations := []Declaration{
		{1, "Reiskosten", "Reiskosten naar de klant"},
		{2, "Lunch", "Lunch met de klant"},
	}

	js, err := json.Marshal(declarations)

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
	// Get request url parameters
	vars := mux.Vars(r)
	varx, err := strconv.Atoi(vars["id"])

	declaration := Declaration{varx, "Reiskosten", "Reiskosten naar de klant"}

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
