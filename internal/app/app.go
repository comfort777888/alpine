package app

import (
	"log"
	"net/http"

	"ascii-art-web/internal/handlers"
)

const (
	port = ":8080"
)

func App() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GETHandler)
	mux.HandleFunc("/ascii-art", handlers.POSTHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	log.Println("Starting server on :8080")
	return http.ListenAndServe(port, mux)
}
