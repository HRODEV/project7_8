package project7_8

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/gorilla/mux"
)

func ReceiptIdGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Get receipt with the specified ID
	var Receipt models.Receipt
	dbActions.GetReceiptById(uint(id), &Receipt, utils.db)

	return &Receipt
}

func ReceiptIdImageGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Get request url parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Get receipt with the specified ID
	var Receipt models.Receipt
	dbActions.GetReceiptById(uint(id), &Receipt, utils.db)

	img, err := os.Open(Receipt.ImagePath)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)

	return nil
}

func ReceiptPost(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	err := r.ParseMultipartForm(100000)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["image"][0]

	file, err := files.Open()
	defer file.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Send request to Microsoft OCR
	var ocrService = services.OcrService{}
	res, err := ocrService.SendImage(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	ocrString := ocrService.GetBoxRightOfWord("Totaal")
	totalPrice, _ := strconv.ParseFloat(strings.Replace(ocrString, ",", ".", -1), 32)

	// Make sure the upload directory does exists
	if _, err := os.Stat("./declarations_upload"); os.IsNotExist(err) {
		os.Mkdir("./declarations_upload", os.ModePerm)
	}

	// Create a empty file and write the uploaded image
	dst, err := os.Create("./declarations_upload/" + files.Filename)
	defer dst.Close()

	// Save the file
	io.Copy(dst, file)

	// Save receipt in the database
	ocrData, _ := json.Marshal(res)
	receipt := models.Receipt{ImagePath: "./declarations_upload/" + files.Filename, Data: string(ocrData)}
	dbActions.CreateReceipt(&receipt, utils.db)

	return &models.Declaration{TotalPrice: float32(totalPrice), ReceiptID: receipt.ID}
}
