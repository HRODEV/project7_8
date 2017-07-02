package project7_8

import (
	"bytes"
	"encoding/json"
	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	http.Error(w, "Not implemented yet", http.StatusNotImplemented)
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

	// Find the price
	ocrResult := ocrService.GetWordsRightOfRgx([][]string{{`\A(?:tota)(?:a)?(?:l)?\z`, `\d+(\.\s?|,\s?|[^a-zA-Z\d])\d{2}`}})

	var totalPrice = 0.0

	if len(ocrResult) > 0 {
		totalPrice, _ = strconv.ParseFloat(strings.Replace(ocrResult[0], ",", ".", -1), 32)
	} else {
		totalPrice = 0
	}

	// Save receipt in the database
	ocrData, _ := json.Marshal(res)
	receipt := models.Receipt{ID: 0, ImagePath: "./declarations_upload/" + files[0].Filename, Data: string(ocrData)}
	dbActions.CreateReceipt(&receipt, utils.db)

	return &models.Declaration{TotalPrice: float32(totalPrice), ReceiptID: receipt.ID}
}
