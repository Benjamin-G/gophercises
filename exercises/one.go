package exercises

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func One() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken`

	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
