package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rafael84/learning/goroutines-and-channels/api"
	"github.com/rafael84/learning/goroutines-and-channels/client"
)

func main() {
	r := mux.NewRouter()

	dailyExports := &client.DailyExports{API: api.NewHTTP("http://localhost:8080")}
	r.Handle("/daily-exports/discover-new-files", dailyExports).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("third-party-client service listening at port 8081")
	log.Fatal(server.ListenAndServe())
}
