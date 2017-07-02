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
	"log"
	"net/http"
	"regexp"
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

// Loop
func (OcrService *OcrService) loopAccrossWords(action func(word *models.OcrWord)) {
	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				action(&word)
			}
		}
	}
}

// Get the the words right of the given regexes
func (OcrService *OcrService) GetWordsRightOfRgx(rgxToSearch [][]string) []string {
	tmpResults := [][]string{}

	// Loop trough all words and save the results of each regex in it's own array [[result1], [result2], etc]
	OcrService.loopAccrossWords(func(word *models.OcrWord) {
		for i, rgxs := range rgxToSearch {
			var rgxResult = regexp.MustCompile(rgxs[0]).FindString(strings.ToLower(word.Text))

			if rgxResult == "" {
				continue
			}

			foundBoundingBox := OcrService.explodeBoundingBox(word.BoundingBox)
			results := OcrService.findWordsRightOfBoudingBox(foundBoundingBox)

			// Since the search regexes are applied to each word, there are more than 1 matches, so we append
			// the result if there is allread a value, otherwise create a new array
			if len(tmpResults) == 0 || len(tmpResults) < i-1 {
				tmpResults = append(tmpResults, []string{})
			}

			tmpResults[i] = append(tmpResults[i], results...)
		}
	})

	// Loop trough the tmpResults and apply the result regex to get the final results
	var finalResults = []string{}

	for i, tmpResult := range tmpResults {
		var resultRgx = regexp.MustCompile(rgxToSearch[i][1])

		// Contact the results to apply the resultrgx
		var contactinatedResult = ""

		for _, result := range tmpResult {
			contactinatedResult += "." + result
		}

		log.Print(contactinatedResult)

		finalResults = append(finalResults, resultRgx.FindString(contactinatedResult))
	}

	return finalResults
}

// Get all words right of the given boundingBox
func (OcrService *OcrService) findWordsRightOfBoudingBox(box models.OcrBoundingBox) []string {
	results := []string{}

	for _, region := range OcrService.OcrData.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				if !OcrService.intersectWithBoundingBox(box, OcrService.explodeBoundingBox(word.BoundingBox)) {
					continue
				}

				results = append(results, word.Text)
			}
		}
	}

	return results
}

// Explode a boundingbox to a `OcrBoundingBox` struct
func (OcrService *OcrService) explodeBoundingBox(box string) models.OcrBoundingBox {
	splittedBox := strings.Split(box, ",")

	x, _ := strconv.Atoi(splittedBox[0])
	y, _ := strconv.Atoi(splittedBox[1])
	width, _ := strconv.Atoi(splittedBox[2])
	height, _ := strconv.Atoi(splittedBox[3])

	return models.OcrBoundingBox{x, y, width, height}
}

// Check if the given box intersects with b
func (OcrService *OcrService) intersectWithBoundingBox(knownBox models.OcrBoundingBox, box models.OcrBoundingBox) bool {
	if knownBox == box {
		return false
	}

	// Calculate the middle of the knownbox
	middle := knownBox.Y + (knownBox.Height / 2)

	return middle > box.Y && middle < box.Y+box.Height && box.X > knownBox.X+knownBox.Width
}
