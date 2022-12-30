package exercises

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
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

type Data struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
	NewYork struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"new-york"`
	Debate struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"debate"`
	SeanKelly struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"sean-kelly"`
	MarkBates struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"mark-bates"`
	Denver struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"denver"`
	Home struct {
		Title   string        `json:"title"`
		Story   []string      `json:"story"`
		Options []interface{} `json:"options"`
	} `json:"home"`
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./gopher.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("origin: %s\n", payload.Intro.Title)
	log.Printf("origin: %s\n", payload.Intro.Story)

	tmpl := template.Must(template.ParseFiles("./html/layout.html"))
	// data := TodoPageData{
	// 	PageTitle: "My TODO list",
	// 	Todos: []Todo{
	// 		{Title: "Task 1", Done: false},
	// 		{Title: "Task 2", Done: true},
	// 		{Title: "Task 3", Done: true},
	// 	},
	// }
	tmpl.Execute(w, payload.Intro)
}

func Three() {
	fmt.Println("Exercise Three Started...")
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/", exampleHandler)
	http.ListenAndServe(":8080", nil)
}
