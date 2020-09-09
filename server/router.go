package server

import (
	"Receiptify/server/controllers"

	"github.com/gorilla/mux"
)

const port := ":8086"

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("base64string", controllers.VisionHandler).Methods("POST")

	http.ListenAndServe(port, router)
}
