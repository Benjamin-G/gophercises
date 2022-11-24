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
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Usage = func() {
		fmt.Printf("Usage of %s \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	records, err := readData(*filename)

	if err != nil {
		log.Fatal(err)
	}

	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press ENTER to continue")
	reader.ReadString('\n')

	done := false
	time.AfterFunc(time.Duration(*limit)*time.Second, func() {
		fmt.Println("\nPress ENTER to see your results...")
		// loop stops at this point
		done = true
	})

	var correct int
	var completed int

	for i, v := range records {
		if done {
			break
		}
		question, answer := func() (string, string) {
			return v[0], v[1]
		}()
		fmt.Printf("Problem #%d: %s:", i+1, question)
		ans, _ := reader.ReadString('\n')

		if runtime.GOOS == "windows" {
			ans = strings.Replace(ans, "\r\n", "", -1)
		} else {
			ans = strings.Replace(ans, "\n", "", -1)
		}

		if strings.Compare(answer, ans) == 0 {
			correct += 1
		}

		completed = i + 1
	}

	fmt.Printf("You scored %d out of %d", correct, completed)
}

// func game(records [][]string, limit int) (int, int)

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
