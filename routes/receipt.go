package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
)

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
	err := r.ParseMultipartForm(100000)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["image"]

	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//create destination file making sure the path is writeable.
		dst, err := os.Create("/home/niels/" + files[i].Filename)
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
