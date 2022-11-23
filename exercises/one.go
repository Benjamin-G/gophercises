package exercises

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
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

	records, err := readData(*filename)

	if err != nil {
		log.Fatal(err)
	}

	done := false
	time.AfterFunc(12*time.Second, func() {
		// Printed after stated duration
		// by AfterFunc() method is over
		fmt.Println("\n12 seconds over....")
		// loop stops at this point
		done = true
	})

	var correct int
	var completed int

	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	reader := bufio.NewReader(os.Stdin)

	for i, v := range records {
		question, answer := func() (string, string) {
			return v[0], v[1]
		}()
		fmt.Printf("Problem #%d: %s:", i, question)
		ans, _ := reader.ReadString('\n')

		if runtime.GOOS == "windows" {
			ans = strings.Replace(ans, "\r\n", "", -1)
		} else {
			ans = strings.Replace(ans, "\n", "", -1)
		}

		if strings.Compare(answer, ans) == 0 {
			correct += 1
		}

		completed = i

		if done {
			break
		}
	}

	fmt.Printf("You scored %d out of %d", correct, completed)
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
