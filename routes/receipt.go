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
	"log"
	"net/http"
	"os"
	"regexp"
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

	var rgx = regexp.MustCompile(`\d+(\.\s?|,\s?|[^a-zA-Z\d])\d{2}`)

	// @TODO Move the processing of the result to the 'OcrService'
	// @TODO Serach with list of regexes
	ocrResult := ocrService.GetWordsRightOf([]string{"totaal", "totaa", "subtota"})
	log.Print(ocrResult)
	combinedResult := ""

	// @TODO reverse resulte before regex
	for _, result := range ocrResult {
		for _, result2 := range result {
			combinedResult += "." + result2
		}
	}

	log.Print("raw result: " + combinedResult)
	log.Print("regex result: " + strings.Replace(rgx.FindString(combinedResult), ",", ".", -1))

	rgxResult := strings.Replace(rgx.FindString(combinedResult), ",", ".", -1)

	totalPrice, _ := strconv.ParseFloat(rgxResult, 32)

	// Save receipt in the database
	ocrData, _ := json.Marshal(res)
	receipt := models.Receipt{ID: 0, ImagePath: "./declarations_upload/" + files[0].Filename, Data: string(ocrData)}
	dbActions.CreateReceipt(&receipt, utils.db)

	return &models.Declaration{TotalPrice: float32(totalPrice), ReceiptID: receipt.ID}
}
