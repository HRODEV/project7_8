package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HRODEV/project7_8/models"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
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

func (OcrService *OcrService) explodeBoundingBox(box string) models.OcrBoundingBox {
	splittedBox := strings.Split(box, ",")

	x, _ := strconv.Atoi(splittedBox[0])
	y, _ := strconv.Atoi(splittedBox[1])
	width, _ := strconv.Atoi(splittedBox[2])
	height, _ := strconv.Atoi(splittedBox[3])

	return models.OcrBoundingBox{x, y, width, height}
}

func (OcrService *OcrService) GetBoxRightOfWord(wordToSearch string) string {
	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				if !strings.Contains(word.Text, wordToSearch) {
					continue
				}

				foundBoundingBox := OcrService.explodeBoundingBox(word.BoundingBox)

				for _, word := range line.Words {
					wordBox := OcrService.explodeBoundingBox(word.BoundingBox)

					if wordBox != foundBoundingBox {
						if foundBoundingBox.Y <= wordBox.Y+20 && foundBoundingBox.Y+foundBoundingBox.Height >= (wordBox.Y+wordBox.Height)-20 {
							return word.Text
						}
					}
				}
			}
		}
	}

	return ""
}
