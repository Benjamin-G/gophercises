package fundementals

import "fmt"

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
