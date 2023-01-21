package fundementals

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	if false {
		t.Fatal("expected false, got true")
	}
}

func Test_GetAlphabet_Errorf(t *testing.T) {
	key := "US"
	alpha, err := getAlphabet(key)
	if err != nil {
		// t.Errorf("could not find alphabet %s", key)
		t.Fatalf("could not find alphabet %s", key)
	}
	act := alpha[12]
	exp := "M"
	if act != exp {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func Test_AddTen(t *testing.T) {
	act := addTen(1)
	exp := 11
	if act != exp {
		t.Fatalf("expected %d, got %d", exp, act)
	}
}

func Test_Chapters(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	chapterTwo()
	chapterThree()
	chapterFour()
	// chapterFive()
	chapterSix()
}

// func Test_TableDrivenTests_Anatomy(t *testing.T) {
// 	t.Parallel()
// 	// any setup code common to all
// 	// test cases goes here
// 	// create a slice of anonymous structs
// 	// initalize the slice with each of the
// 	// desired test cases and assign it to the
// 	// variable 'tcs'. tcs stands for "test cases"
// 	tcs := []struct {
// 		// fields needed for each test case
// 	}{
// 		// { tests cases go here },
// 		// { tests cases go here },
// 	}
// 	for _, tc := range tcs { // tc stands for "test case"
// 		// loop through each test case
// 		// and make the necessary assertions
// 		// for that test case
// 	}
// }

func (s *Store) All(tn string) (Model, error) {
	db := s.data
	// if the underlying data is nil, return an error
	if db == nil {
		return nil, fmt.Errorf("no data")
	}
	// check to make sure table exists
	mods, ok := db[tn]
	// if table doesn't exist, return an error
	if !ok {
		return nil, fmt.Errorf("table %s not found", tn)
	}
	// return the slice of models
	return mods, nil
}

func Test_Store_All_Errors(t *testing.T) {
	t.Parallel()
	tn := "users"
	tcs := []struct {
		name  string
		store *Store
		exp   error
	}{
		// {name: "no data", store: noData(t), exp: ErrNoData(tn)},
		// {name: "with data, no users", store: withData(t), exp: ErrTableNotFound{}},
		{name: "with users", store: withUsers(t), exp: nil},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.store.All(tn)
			ok := errors.Is(err, tc.exp)
			if !ok {
				t.Fatalf("expected error %v, got %v", tc.exp, err)
			}
		})
	}
}

type ErrTableNotFound struct {
	Table      string
	OccurredAt time.Time
}

func (e ErrTableNotFound) Error() string {
	return fmt.Sprintf("[%s] table not found %s", e.OccurredAt, e.Table)
}

func noData(t testing.TB) *Store {
	return &Store{}
}
func withData(t testing.TB) *Store {
	return &Store{
		data: map[string]Model{},
	}
}

func withUsers(t testing.TB) *Store {
	t.Helper()
	users := Model{
		"id": 1,
		// {"id": 2, "name": "Jane"},
	}
	t.Cleanup(func() {
		t.Log("cleaning up users", users)
	})
	return &Store{
		data: map[string]Model{
			"users": users,
		},
	}
}
