package main

import (
	"github.com/DangerOnTheRanger/seenit"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	dbEnvVar = "SEENIT_DATABASE_PATH"
	defaultDbPath = "database.db"
)

func main() {
	dbPath := os.Getenv(dbEnvVar)
	if dbPath == "" {
		dbPath = defaultDbPath
	}
	db, err := seenit.NewBoltDatabase(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("database found at %s", dbPath)
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", seenit.ServeLanding)
	router.HandleFunc("/upload", seenit.ServeUpload)
	router.HandleFunc("/result", seenit.BindHandler(seenit.ServeResult, db))

	log.Fatal(http.ListenAndServe(":8080", router))
}
