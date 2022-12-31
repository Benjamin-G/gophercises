package exercises

import (
	"flag"
	"fmt"
	"gophercises/urlshort"
	"log"
	"net/http"
	"os"
)

var fileName *string

// func init() {
// 	fileName = flag.String("yml", "exe2.yml", "a yml file of mapped paths to url")

// 	flag.Usage = func() {
// 		fmt.Printf("Usage of %s \n", os.Args[0])
// 		flag.PrintDefaults()
// 	}

// 	flag.Parse()
// }

func readDataString(fileName string) ([]byte, error) {
	f, err := os.ReadFile(fileName)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	if err != nil {
		return []byte(yaml), err
	}

	return f, nil
}

func Two() {
	fileName = flag.String("yml", "exe2.yml", "a yml file of mapped paths to url")

	flag.Usage = func() {
		fmt.Printf("Usage of %s \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	// filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	fmt.Println(fileName)
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yml, err := readDataString(*fileName)
	if err != nil {
		log.Println("Error readDataString:", err)
	}

	yamlHandler, err := urlshort.YAMLHandler(yml, mapHandler)
	if err != nil {
		log.Println("Error YAMLHandler:", err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
