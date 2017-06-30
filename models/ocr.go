package models

type Ocr struct {
	Language    string       `json:"language"`
	TextAngle   float32      `json:"textAngle"`
	Orientation string       `json:"orientation"`
	Regions     []OcrRegions `json:"regions"`
}

type OcrRegions struct {
	BoundingBox string     `json:"boundingBox"`
	Lines       []OcrLines `json:"lines"`
}

type OcrLines struct {
	BoundingBox string    `json:"boundingBox"`
	Words       []OcrWord `json:"words"`
}

type OcrWord struct {
	BoundingBox string `json:"boundingBox"`
	Text        string `json:"text"`
}

type OcrBoundingBox struct {
	X      int
	Y      int
	Width  int
	Height int
}
