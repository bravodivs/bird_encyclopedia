package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", Index).Methods("GET")
	r.HandleFunc("/birds", getBird).Methods("GET")
	r.HandleFunc("/birds", createBird).Methods("POST")

	// adding a static asset:
	// 1. create static asset
	// 2. modify the router to serve static assets
	staticFileDirectory := http.Dir("./static")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/static/" prefix when looking for files.
	// For example, if we type "/static/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./static/static/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/static/", instead of the absolute route itself
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}
