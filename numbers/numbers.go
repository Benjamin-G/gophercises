package numbers

import (
	"fmt"
	"gophercises/database"
)

func Numbers() {
	fmt.Println("Numbers Starting...")
	db, err := database.SetupNumbersDBConnection()
	must(err)
	defer db.Close()

	err = database.CreatePhoneNumbersTable(db)
	must(err)

	numbersToAdd := []string{"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}

	for _, num := range numbersToAdd {
		err = database.CreatePhoneNumber(db, num)
		must(err)
	}

	// // var p = database.PhoneRecord{Id: 4, Number: "(123) 456 7892"}
	// database.UpdatePhoneNumber(db, database.PhoneRecord{Id: 4, Number: "(123) 456 7892"})

	done := make(chan interface{})
	defer close(done)
	database.CleanPhoneNumbers(db, done)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
