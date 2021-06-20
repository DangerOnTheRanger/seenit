package main

import (
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ServeLanding)
	router.HandleFunc("/upload", ServeUpload)
	router.HandleFunc("/result", ServeResult)
}
