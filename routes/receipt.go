package project7_8

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"bytes"
	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/gorilla/mux"
	"io/ioutil"
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

	if Receipt.ID == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return nil
	}

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

	// Make sure the image exists
	if _, err := os.Stat(Receipt.ImagePath); os.IsNotExist(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	img, err := os.Open(Receipt.ImagePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
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
	files := m.File["image"]

	if len(files) == 0 || len(files) > 1 {
		http.Error(w, "No file was found in the 'image' header or multiple files are send", http.StatusInternalServerError)
		return nil
	}

	file, err := files[0].Open()
	defer file.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Make sure the upload directory does exists
	if _, err := os.Stat("./declarations_upload"); os.IsNotExist(err) {
		os.Mkdir("./declarations_upload", os.ModePerm)
	}

	// Create a empty file and write the uploaded image
	dst, err := os.Create("./declarations_upload/" + files[0].Filename)
	defer dst.Close()

	// Convert file to reader
	imageData, _ := ioutil.ReadAll(file)

	// Save the file
	io.Copy(dst, bytes.NewReader(imageData))

	// Send request to microsoft
	ocrService := services.OcrService{}
	res, err := ocrService.SendImage(bytes.NewReader(imageData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	ocrResult := ocrService.GetBoxRightOfWord("totaal")
	totalPrice, _ := strconv.ParseFloat(strings.Replace(ocrResult, ",", ".", -1), 32)

	// Save receipt in the database
	ocrData, _ := json.Marshal(res)
	receipt := models.Receipt{ID: 0, ImagePath: "./declarations_upload/" + files[0].Filename, Data: string(ocrData)}
	dbActions.CreateReceipt(&receipt, utils.db)

	return &models.Declaration{TotalPrice: float32(totalPrice), ReceiptID: receipt.ID}
}
