package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/HRODEV/project7_8/models"
)

type OcrService struct {
	OcrData models.Ocr
}

func (OcrService *OcrService) SendImage(image multipart.File) (*models.Ocr, error) {
	// Create a empty buffer and client
	body := &bytes.Buffer{}
	client := &http.Client{}

	// Write the multipart formdata
	multipartWriter := multipart.NewWriter(body)
	fw, _ := multipartWriter.CreateFormFile("image", "image.jpg")
	io.Copy(fw, image)
	multipartWriter.Close()

	// Construct and send the request
	req, _ := http.NewRequest("POST", "https://westus.api.cognitive.microsoft.com/vision/v1.0/ocr", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Ocp-Apim-Subscription-Key", "4187c1df33514aa7b412a1eefcacbde4")

	res, err := client.Do(req)
	responseBody, _ := ioutil.ReadAll(res.Body)

	if err != nil || res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s; message:%s", res.Status, responseBody)

		return nil, err
	}

	var ocr *models.Ocr
	json.Unmarshal(responseBody, &ocr)

	OcrService.OcrData = *ocr

	return ocr, nil
}

func (OcrService *OcrService) GetBoxRightOfWord(wordToSearch string) string {
	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				if !strings.Contains(strings.ToLower(word.Text), strings.ToLower(wordToSearch)) {
					continue
				}

				foundBoundingBox := OcrService.explodeBoundingBox(word.BoundingBox)

				return OcrService.findWordsInBoudingBox(foundBoundingBox)
			}
		}
	}

	return ""
}

func (OcrService *OcrService) findWordsInBoudingBox(box models.OcrBoundingBox) string {
	for _, region := range OcrService.OcrData.Regions {
		if !OcrService.isInBoundingBox(OcrService.explodeBoundingBox(region.BoundingBox), box) {
			continue
		}

		for _, line := range region.Lines {
			if !OcrService.isInBoundingBox(OcrService.explodeBoundingBox(line.BoundingBox), box) {
				continue
			}
			for _, word := range line.Words {
				if !OcrService.isInBoundingBox(OcrService.explodeBoundingBox(word.BoundingBox), box) {
					continue
				}

				return word.Text
			}
		}
	}

	return ""
}

func (OcrService *OcrService) explodeBoundingBox(box string) models.OcrBoundingBox {
	splittedBox := strings.Split(box, ",")

	x, _ := strconv.Atoi(splittedBox[0])
	y, _ := strconv.Atoi(splittedBox[1])
	width, _ := strconv.Atoi(splittedBox[2])
	height, _ := strconv.Atoi(splittedBox[3])

	return models.OcrBoundingBox{x, y, width, height}
}

// Check if the given box lies in the boundings of b
func (OcrService *OcrService) isInBoundingBox(b models.OcrBoundingBox, box models.OcrBoundingBox) bool {
	return b.Y <= (box.Y+20) && (b.Y+b.Height) >= (box.Y+box.Height-20)
}
