package exercises

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./html/layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func Three() {
	fmt.Println("Exercise Three Started...")
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/", exampleHandler)
	http.ListenAndServe(":8080", nil)
}
