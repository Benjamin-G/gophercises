package fundementals

import (
	"fmt"
	"log"
	"os"
)

func FilesMain() {
	files, err := os.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("->", file.Name())
			continue
		}
		fmt.Println(file.Name())
	}
}
