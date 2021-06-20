package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DangerOnTheRanger/seenit"
)

func main() {
	db := seenit.MockDatabase{Buckets: make(map[string]seenit.MockBucket, 0)}
	
	router := mux.NewRouter()
	router.HandleFunc("/", seenit.ServeLanding)
	router.HandleFunc("/upload", seenit.ServeUpload)
	router.HandleFunc("/result", seenit.BindHandler(seenit.ServeResult, &db))

	http.ListenAndServe(":8080", router)
}
