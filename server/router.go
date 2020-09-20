package server

import (
	"Receiptify/server/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8086"

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("api/v1/base64string", controllers.VisionHandler).Methods("POST")

	// need to account for all the statics
	fs := http.FileServer(http.Dir("./client/receiptify/dist"))
	router.PathPrefix("/js").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/css").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/img").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", fs))

	router.NotFoundHandler = http.HandlerFunc(controllers.IndexRoute)

	http.ListenAndServe(port, router)
}
