package exercises

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type PageObject struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func unmarshalPageObject(data []byte) map[string]PageObject {
	var z map[string]PageObject

	if err := json.Unmarshal(data, &z); err != nil {
		log.Fatal("Error when Unmarshaling json: ", err)
		return nil
	}

	return z
}

func cyoa(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./gopher.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	pages := unmarshalPageObject(content)

	query := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/"))
	log.Printf("query: %s\n", query)

	page := pages["intro"]
	for k := range pages {
		if k == query {
			page = pages[k]
			break
		}
	}

	tmpl := template.Must(template.ParseFiles("./html/layout.html"))
	tmpl.Execute(w, page)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func Three() {
	fmt.Println("Exercise Three Started...")
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/", cyoa)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":8080", nil)
}
