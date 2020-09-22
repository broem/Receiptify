package server

import (
	"Receiptify/server/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const port = ":8086"

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/base64string", controllers.VisionHandler).Methods("POST")

	// need to account for all the static
	fs := http.FileServer(http.Dir("./client/receiptify/dist"))
	router.PathPrefix("/js").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/css").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/img").Handler(http.StripPrefix("/", fs))
	router.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", fs))

	router.NotFoundHandler = http.HandlerFunc(controllers.IndexRoute)

	router.Use(Middleware)

	log.Printf("Now listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func UseJson(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func UseCors(h http.Handler) http.Handler {
	return cors.Default().Handler(h)
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("From: %s -> %s: %s", r.Host, r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func Middleware(h http.Handler) http.Handler {
	h = UseJson(h)
	h = UseCors(h)
	return Logger(h)
}
