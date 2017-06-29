package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	var Receipt models.Receipt

	if db.Where("ID = ?", id).Find(&Receipt).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	} else {
		enc := json.NewEncoder(w)
		enc.Encode(Receipt)
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

		// Send request to Microsoft OCR
		var ocrService = services.OcrService{}
		res, err := ocrService.SendImage(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ocrString := ocrService.GetBoxRightOfWord("Totaal")
		totalPrice, _ := strconv.ParseFloat(strings.Replace(ocrString, ",", ".", -1), 32)

		// Make sure the upload directory does exists
		if _, err := os.Stat("./declarations_upload"); os.IsNotExist(err) {
			os.Mkdir("./declarations_upload", os.ModePerm)
		}

		// Create a empty file and write the uploaded image
		dst, err := os.Create("./declarations_upload/" + files[i].Filename)
		defer dst.Close()

		// Save the file
		io.Copy(dst, file)

		// Save receipt in the database
		ocrData, _ := json.Marshal(res)
		receipt := models.Receipt{ID: 0, ImagePath: "./declarations_upload/" + files[i].Filename, Data: string(ocrData)}
		db.LogMode(true)
		db.Create(&receipt)

		enc := json.NewEncoder(w)
		enc.Encode(&models.Declaration{TotalPrice: float32(totalPrice), ReceiptID: receipt.ID, Date: time.Now().Format(time.RFC3339)})
	}
}
