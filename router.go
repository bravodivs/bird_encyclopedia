package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", Index).Methods("GET")
	r.HandleFunc("/birds", getBird).Methods("GET")
	r.HandleFunc("/birds", createBird).Methods("POST")

	return r
}
