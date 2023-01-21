package fundementals

import (
	"fmt"
	"math/rand"
)

func getAlphabet(key string) ([]string, error) {
	az := []string{"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z"}
	m := map[string][]string{
		"US": az,
		"UK": az,
	}
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("no key found %s", key)
}

func addTen(i int) int {
	return i + 10
}

// Model is a key/value pair representing a model in the store.
// e.g. {"id": 1, "name": "bob"}
type Model map[string]any

// Store is a table based key/value store.
type Store struct {
	data map[string]Model
}

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func createMatrix512(rows int) [][512]int64 {
	var res [][512]int64

	for i := 0; i < rows; i++ {
		res = append(res, [512]int64{}) // append to slice an empty array
		for j := 0; j < 512; j++ {
			res[i][j] = rand.Int63()
		}
	}

	return res
}

func createMatrix513(rows int) [][513]int64 {
	var res [][513]int64

	for i := 0; i < rows; i++ {
		res = append(res, [513]int64{}) // append to slice an empty array
		for j := 0; j < 513; j++ {
			res[i][j] = rand.Int63()
		}
	}

	return res
}

func calculateSum512(s [][512]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ { // Iterate over each row
		for j := 0; j < 8; j++ { // Iterate over the first eight columns
			sum += s[i][j] // Increment sum
		}
	}
	return sum
}

func calculateSum513(s [][513]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ { // Iterate over each row
		for j := 0; j < 8; j++ { // Iterate over the first eight columns
			sum += s[i][j] // Increment sum
		}
	}
	return sum
}
