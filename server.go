package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
}

func todos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/todos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	err = t.Execute(w, nil)

}

const port = "8080"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos", todos)
	fmt.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
