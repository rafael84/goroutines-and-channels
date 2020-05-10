package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rafael84/learning/goroutines-and-channels/server"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/daily-exports", server.DailyExportsHandler)
	r.PathPrefix(server.FilesPath).Handler(http.StripPrefix(server.FilesPath, http.FileServer(http.Dir(server.FilesDir))))

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Println("third-party-server listening at port 8080")
	log.Fatal(server.ListenAndServe())
}
