package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// var db *sql.DB

// type phone struct {
// 	id     int
// 	number string
// }

const (
	host   = "localhost"
	port   = 5432
	user   = "ben"
	dbname = "go_lang_exercises"
)

// func init() {
// 	fmt.Println("Setting db connection")
// 	password := os.Getenv("PGSQLPW")
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)

// 	var openDb error
// 	db, openDb = sql.Open("postgres", psqlInfo)
// 	must(openDb)
// }

func SetupNumbersDBConnection() (*sql.DB, error) {
	fmt.Println("Setting db connection")
	password := os.Getenv("PGSQLPW")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}

func CreatePhoneNumbersTable(db *sql.DB) error {
	statement := `
    CREATE TABLE IF NOT EXISTS public.phone_numbers (
      id SERIAL,
      value VARCHAR(255)
    )`
	_, err := db.Exec(statement)

	return err
}

func CreatePhoneNumber(db *sql.DB, num string) error {
	statement := `INSERT INTO public.phone_numbers ( value ) values ($1) RETURNING id`

	var id int
	err := db.QueryRow(statement, num).Scan(&id)
	fmt.Println(id)

	// No need for return, use execute statement
	// statement := `INSERT INTO public.phone_numbers ( value ) values ($1)`
	// res, err := db.Exec(statement, num)
	// id, _ := res.LastInsertId()
	// fmt.Println(id)

	return err
}
