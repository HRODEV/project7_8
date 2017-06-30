package services

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"

	"fmt"
	"github.com/HRODEV/project7_8/models"
	"net/http"
)

type OcrService struct {
	OcrData models.Ocr
}

func (OcrService *OcrService) SendImage(image io.Reader) (*models.Ocr, error) {
	// Create a empty buffer and client
	var body bytes.Buffer
	client := &http.Client{}

	// Write the multipart formdata
	multipartWriter := multipart.NewWriter(&body)
	fw, _ := multipartWriter.CreateFormFile("image", "image.png")
	io.Copy(fw, image)
	multipartWriter.Close()

	// Construct and send the request
	req, _ := http.NewRequest("POST", "https://westus.api.cognitive.microsoft.com/vision/v1.0/ocr", &body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Ocp-Apim-Subscription-Key", "4187c1df33514aa7b412a1eefcacbde4")

	res, err := client.Do(req)
	responseBody, _ := ioutil.ReadAll(res.Body)

	if err != nil || res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s; message:%s", res.Status, responseBody)

		return nil, err
	}

	// Format response as OCR model
	var ocr *models.Ocr
	json.Unmarshal(responseBody, &ocr)

	OcrService.OcrData = *ocr

	return ocr, nil
}

func (OcrService *OcrService) GetWordsRightOf(wordsToSearch []string) [][]string {
	results := [][]string{}

	OcrService.loop(func(word *models.OcrWord) {
		for _, search := range wordsToSearch {
			if strings.Contains(strings.ToLower(word.Text), strings.ToLower(search)) {
				foundBoundingBox := OcrService.explodeBoundingBox(word.BoundingBox)

				results = append(results, OcrService.findWordsInBoudingBox(foundBoundingBox))
			}
		}
	})

	return results
}

func (OcrService *OcrService) loop(action func(word *models.OcrWord)) {
	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				action(&word)
			}
		}
	}
}

func (OcrService *OcrService) findWordsInBoudingBox(box models.OcrBoundingBox) []string {
	results := []string{}

	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				if !OcrService.intersectWithBoundingBox(OcrService.explodeBoundingBox(word.BoundingBox), box) {
					continue
				}

				results = append(results, word.Text)
			}
		}
	}

	return results
}

func (OcrService *OcrService) explodeBoundingBox(box string) models.OcrBoundingBox {
	splittedBox := strings.Split(box, ",")

	x, _ := strconv.Atoi(splittedBox[0])
	y, _ := strconv.Atoi(splittedBox[1])
	width, _ := strconv.Atoi(splittedBox[2])
	height, _ := strconv.Atoi(splittedBox[3])

	return models.OcrBoundingBox{x, y, width, height}
}

// Check if the given box lies intersects with b
func (OcrService *OcrService) intersectWithBoundingBox(b models.OcrBoundingBox, box models.OcrBoundingBox) bool {
	if b == box {
		return false
	}

	middle := box.Y + (box.Height / 2)

	// True when `middle` intersects with box b; only searches on the right of b
	return middle > b.Y && middle < b.Y+b.Height && box.X < b.X+b.Width
}
