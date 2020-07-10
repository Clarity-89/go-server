package main

import (
	"fmt"
	"go-server/models"
	"go-server/services/storage"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = "8080"

type Todo struct {
	Title   string
	Content string
}

type PageContext struct {
	Title string
	Todos []Todo
}

var todos []Todo

var s = storage.Init()

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	err = t.Execute(w, nil)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	pageContext := PageContext{
		Todos: todos,
		Title: "Todos",
	}

	t, err := template.ParseFiles("templates/todos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	err = t.Execute(w, pageContext)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Form parsing error:", err)
	}

	todo := models.TodoDTO{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
		DueDate: time.Now(),
	}

	err = s.SaveTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Data saving error:", err)
		return
	}

	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", listTodos)
	http.HandleFunc("/add-todo", addTodo)
	fmt.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
