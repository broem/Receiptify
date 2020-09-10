package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/otiai10/gosseract/v2"
)

// Vision returns the tesseracted information
func VisionHandler(w http.ResponseWriter, req *http.Request) {
	visions := new(VisionRequest)

	err := json.NewDecoder(req.Body).Decode(visions)
	if err != nil {

	}

	tempF, err := ioutil.TempFile("", "")
	if err != nil {

	}
	defer func() {
		tempF.Close()
		os.Remove(tempF.Name())
	}()

	visions.Base64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(visions.Base64, "")
	b, err := base64.StdEncoding.DecodeString(visions.Base64)
	if err != nil {

	}
	tempF.Write(b)

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(tempF.Name())

	translated, err := client.Text()
	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(translated))

}
