package main

import (
	"fmt"
	"net/http"

	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
}

func todos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Todos")
}

const port = "8080"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos", todos)
	fmt.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
