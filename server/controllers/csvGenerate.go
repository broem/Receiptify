package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GenerateCSV(w http.ResponseWriter, req *http.Request) {
	// take in the raw text, convert white splits to csv
	raw := new(CsvRequest)

	err := json.NewDecoder(req.Body).Decode(raw)
	if err != nil {
		log.Printf("Unable to decode csvreq")
	}

	// should probably map it to a receipt format
	csvified := GenerateCsvFromRaw(raw.RawText)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(csvified)
}

func GenerateCsvFromRaw(text string) string {
	split := strings.Fields(text)
	return strings.Join(split, ",")
}
