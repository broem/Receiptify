package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/otiai10/gosseract/v2"
)

// VisionHandler ... Vision returns the tesseracted information
func VisionHandler(w http.ResponseWriter, req *http.Request) {
	visions := new(VisionRequest)

	err := json.NewDecoder(req.Body).Decode(visions)
	if err != nil {

	}

	translated, err := getRawText(visions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(translated)
}

// VisionCSV ... returns the tesseracted information parsed into a neat csv
func VisionCSV(w http.ResponseWriter, req *http.Request) {

}

func getRawText(v VisionRequest) string, error {

	tempF, err := ioutil.TempFile("", "")
	if err != nil {
		log.Printf("Unable to create temp file")
		return "", err
	}
	defer func() {
		tempF.Close()
		os.Remove(tempF.Name())
	}()

	v.Base64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(v.Base64, "")
	b, err := base64.StdEncoding.DecodeString(v.Base64)
	if err != nil {
		log.Printf("Unable to decode base64")
		return "Unable to decode", err
	}
	tempF.Write(b)

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(tempF.Name())

	translated, err := client.Text()
	if err != nil {
		log.Printf("Unable to get text from image!")
		return "Unable to get text from image!", err
	}

	return translated, err
}
