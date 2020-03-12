package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"startyourday/model"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	data := model.HomePageData{
		PageTitle: "My TODO list",
		Todos: []model.Todo{
			{Name: "Task 1", Description: "do this and that to fulfill task 1"},
			{Name: "Task 2", Description: "do the other and this to fulfill task 2"},
			{Name: "Task 3", Description: "do something to do this task 3"},
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
