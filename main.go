package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello bird")

	// will give health(/health), get birds(/birds, get) and create birds(/birds, post)
	// /static/ redirects to home page
	r := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
