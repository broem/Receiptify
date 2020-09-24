package controllers

type VisionRequest struct {
	Base64 string `json:"base64"`
}

type CsvRequest struct {
	RawText string `json:"rawText"`
}
