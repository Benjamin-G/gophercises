package exercises

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

func One() {
	fmt.Println(flag.Arg(0))

	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")

	fmt.Println(*filename)

	flag.Usage = func() {
		fmt.Printf("Usage of %s \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Must parse the flag arguments before reference
	fmt.Println(*filename)

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF

		fmt.Println(strings.EqualFold("hi", text))

		if runtime.GOOS == "windows" {
			text = strings.Replace(text, "\r\n", "", -1)
		} else {
			text = strings.Replace(text, "\n", "", -1)
		}

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		if strings.Compare("exit", text) == 0 {
			break
		}
	}
	fmt.Println("Bye!")
}
