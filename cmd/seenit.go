package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DangerOnTheRanger/seenit"
)

const (
	dbPath = "database.db"
)

func main() {
	//db := seenit.MockDatabase{Buckets: make(map[string]seenit.MockBucket, 0)}
	db, err := seenit.NewBoltDatabase(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", seenit.ServeLanding)
	router.HandleFunc("/upload", seenit.ServeUpload)
	router.HandleFunc("/result", seenit.BindHandler(seenit.ServeResult, db))

	log.Fatal(http.ListenAndServe(":8080", router))
}
