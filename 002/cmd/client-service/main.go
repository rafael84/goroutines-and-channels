package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/rafael84/goroutines-and-channels/002/api"
	"github.com/rafael84/goroutines-and-channels/002/client"
)

func main() {
	r := mux.NewRouter()

	dailyExports := &client.DailyExports{API: api.NewHTTP("http://localhost:8080")}
	r.Handle("/daily-exports/discover-new-files", dailyExports).Methods("POST")

	server := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "0.0.0.0:8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("third-party-client service listening at port 8081")
	log.Fatal(server.ListenAndServe())
}
