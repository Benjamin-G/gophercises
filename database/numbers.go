package database

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"sync"

	_ "github.com/lib/pq"
)

// Can use a global variable however dependencies injection is preferred
// var db *sql.DB

// func init() {
// 	fmt.Println("Setting db connection")
// 	password := os.Getenv("PGSQLPW")
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)

// 	var openDb error
// 	db, openDb = sql.Open("postgres", psqlInfo)
// 	must(openDb)
// }

func init() {
	fmt.Println("Setting db connection")
}

const (
	host   = "localhost"
	port   = 5432
	user   = "ben"
	dbname = "go_lang_exercises"
)

// To import all need capitalization
type PhoneRecord struct {
	Id     int
	Number string
}

func SetupNumbersDBConnection() (*sql.DB, error) {
	fmt.Println("Setting db connection...")
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
	// If I need a return value
	// statement := `INSERT INTO public.phone_numbers ( value ) values ($1) RETURNING id`

	// var id int
	// err := db.QueryRow(statement, num).Scan(&id)
	// fmt.Println(id)

	// No need for return, use execute statement
	statement := `INSERT INTO public.phone_numbers ( value ) values ($1)`
	_, err := db.Exec(statement, num)
	// id, _ := res.LastInsertId()
	// fmt.Println(id)

	return err
}

func UpdatePhoneNumber(db *sql.DB, p PhoneRecord) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.Id, p.Number)
	return err
}

func CleanPhoneNumbers(db *sql.DB, done <-chan interface{}) error {
	var wg sync.WaitGroup

	allPhonesNumbers, err := allPhones(db)
	if err != nil {
		return err
	}

	normalize := func(phone string) string {
		var buf bytes.Buffer
		for _, ch := range phone {
			if ch >= '0' && ch <= '9' {
				buf.WriteRune(ch)
			}
		}
		return buf.String()
	}

	updateAndClean := func(record <-chan PhoneRecord) {
		for {
			select {
			case <-done:
				return
			case c := <-record:
				UpdatePhoneNumber(db, PhoneRecord{Id: c.Id, Number: normalize(c.Number)})
				wg.Done()
			}
		}
	}

	var needsCleaned = make(chan PhoneRecord)

	// Set up two goroutines parallel with a shared channel, order does not matter, could scale to as many as needed
	go updateAndClean(needsCleaned)
	go updateAndClean(needsCleaned)

	onlyNumbers := regexp.MustCompile(`^[0-9]*$`)

	for _, record := range allPhonesNumbers {
		// Only need to update rows that need to be cleaned
		if !onlyNumbers.Match([]byte(record.Number)) {
			wg.Add(1)
			needsCleaned <- record
		}
	}

	wg.Wait()
	return err
}

func allPhones(db *sql.DB) ([]PhoneRecord, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []PhoneRecord
	for rows.Next() {
		var p PhoneRecord
		if err := rows.Scan(&p.Id, &p.Number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
