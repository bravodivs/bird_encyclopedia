package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getBird(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}
	bird.Species = "Pigeon2"
	bird.Description = "Common bird found in cities"

	birds = append(birds, bird)

	birdsJsonList, err := json.Marshal(birds)
	fmt.Fprintf(w, "Call to get a bird")
	if err != nil {
		log.Fatal("cant convert birds to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// why not using json encoding?
	// json ecoding shud be used for json. for plain text, Write is fine
	w.Write(birdsJsonList)
}
func createBird(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Call to post a bird")
	bird := Bird{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	birds = append(birds, bird)

	fmt.Print(bird)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/static/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/static/", http.StatusFound)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(bird)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello Bird. Health is fine")
}
