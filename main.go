package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello bird")

	r := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
