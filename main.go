package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	SetupRoutes(r)

	var listenAddr string

	if os.Getenv("PORT") == "" {
		listenAddr = ":8000"
	} else {
		listenAddr = ":" + os.Getenv("PORT")
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
