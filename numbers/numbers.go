package numbers

import (
	"fmt"
	"gophercises/database"
)

func Numbers() {
	fmt.Println("Numbers -")
	db, err := database.SetupNumbersDBConnection()
	must(err)
	defer db.Close()

	err = database.CreatePhoneNumbersTable(db)
	must(err)

	err = database.CreatePhoneNumber(db, "1234567890")
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
