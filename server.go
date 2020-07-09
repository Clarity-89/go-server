package main

import (
	"fmt"
	"net/http"

	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
}

const port = "8000"

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
