package main

import (
	"net/http"
	"html/template"
	"log"
)

var templ *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	
	data := PageData {
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Build a Web App", Done: false},
		},
	}
	
	templ.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	templ = template.Must(template.ParseFiles("templates/index.gohtml"))
	
	
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	
	log.Fatal(http.ListenAndServe(":5100", mux))
}
