package database

import (
	"fmt"
	"os"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "ben"
	dbname = "go_lang_exercises"
)

func init() {
	fmt.Println("Hello database!")
	password := os.Getenv("PGSQLPW")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	fmt.Println(psqlInfo)
}

func Numbers() {
	fmt.Println("Numbers")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
