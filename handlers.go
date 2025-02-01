package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getBird(w http.ResponseWriter, r *http.Request) {
	birdsJsonList, err := json.Marshal(birds)

	if err != nil {
		log.Fatal("cant convert birds to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// why not using json encoding?
	w.Write(birdsJsonList)
}
func createBird(w http.ResponseWriter, r *http.Request) {

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello Bird. Health is fine")
}
