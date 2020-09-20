package controllers

import "net/http"

func IndexRoute(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./client/receiptify/dist"+"/index.html")
}
