package fundementals

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func GetString(key string) (string, error) {
	m := map[string]string{
		"a": "A",
		"b": "B",
	}
	if v, ok := m[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("no key found %s", key)
}

func RunErrorMain() {
	// defer func() {
	// 	if i := recover(); i != nil {
	// 		fmt.Println("oh no, a panic occurred:", i)
	// 	}
	// }()
	// a := []string{}
	// a[42] = "Bring a towel"

	// create a matcher function
	m := func(r rune) bool {
		// simulate doing something bad...
		panic("hahaha")
		// unreachable code
		return false
	}

	// sanitize the string
	_, err := sanitize(m, "go is awesome")
	if err != nil {
		// handle the error
		fmt.Println("Error :", err)
	}

	_, err = sanitize(func(r rune) bool {
		// unreachable code
		return false
	}, "go is awesome")

	if err != nil {
		// handle the error
		fmt.Println("Error :", err)
	}

	// MAPS
	// create a new map variable
	var n map[string]int
	if n == nil {
		// initialize the map
		n = map[string]int{}
	}
	// insert a key-value pair
	n["Amy"] = 27
	// print the map
	fmt.Printf("%+v\n", n)

	// OR
	// initialize a new map
	n2 := map[string]int{}
	// insert a key-value pair
	n2["Amy"] = 27
	// print the map
	fmt.Printf("%+v\n", n2)

	// initialize a new map
	n3 := make(map[string]int)
	// insert a key-value pair
	n3["Amy"] = 27
	// print the map
	fmt.Printf("%+v\n", n3)

	// Pointers
	// to a bytes.Buffer
	bb := &bytes.Buffer{}
	// use the pointer to
	// write data to the buffer
	bb.WriteString("Hello, world!")
	// print the buffer
	fmt.Println(bb.String())

	// Proper interface
	// initialize a stream
	// with STDOUT as the writer
	stream := Stream{
		Writer: os.Stdout,
	}
	if stream.Writer != nil {
		fmt.Fprintf(stream, "Hello Gophers!\n")
	}

	stream2 := Stream{}
	if stream2.Writer != nil {
		fmt.Fprintf(stream, "Hello Gophers!\n")
	}

	// create a buffer
	// to write to
	bb2 := &bytes.Buffer{}
	// data to be written
	data := []byte("Hello, world!")
	// call WriteToFile
	// passing the buffer
	// and the data
	err = WriteToFile(bb2, data)
	// check for errors
	if err != nil {
		fmt.Println("Error :", err)
		// os.Exit(1)
	}

	// Arrays and Slices
	// create a slice
	names := []string{"Kurt", "Janis", "Jimi", "Amy"}
	// find index 42
	s, err := find(names, 42)
	if err != nil {
		fmt.Println("Error :", err)
		// os.Exit(1)
	} else {
		fmt.Println(s)
	}

	err = doSomething(1)
	if err != nil {
		fmt.Println("Error :", err)
	}
}

func find(names []string, index int) (string, error) {
	// check for out of bounds index
	if index >= len(names) {
		return "", fmt.Errorf("out of bounds index %d [%d]", index, len(names))
	}
	// find the name at the index
	s := names[index]
	// return an error if the value is empty
	if len(s) == 0 {
		return s, fmt.Errorf("index %d empty", index)
	}
	// return the name
	return s, nil
}

func WriteToFile(w io.Writer, data []byte) error {
	// assert that w is a file
	f, ok := w.(*os.File)
	// check the assertion was successful
	if !ok {
		return fmt.Errorf("expected *os.File, got %T", w)
	}
	// defer closing the file
	defer f.Close()
	// log the file name
	fmt.Printf("writing to file %s\n", f.Name())
	// write the data
	_, err := f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type Stream struct {
	io.Writer
}

type matcher func(rune) bool

func sanitize(m matcher, s string) (val string, err error) {
	// var val string
	// var err error
	// guard against an invalid matcher that could panic
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("invalid matcher. panic occurred: %v", e)
		}
	}()
	// iterate over the runes in the string
	for _, c := range s {
		// call the matcher function
		// with the rune as the argument
		if m(c) {
			// append '*' to the result
			val = val + "*"
			// continue to the next rune
			continue
		}
		// append the rune to the result
		val = val + string(c)
	}
	// return the sanitized string
	return val, err
}

func doSomething(input int) (err error) {
	// defer a function to recover from the panic
	defer func() {
		p := recover()
		if p == nil {
			// a nil was return, no panic was raised
			// return from the deferred function.
			return
		}
		// check if the recovered value is already an error
		if e, ok := p.(error); ok {
			// assign the recovered error to the err variable
			// outside of the anonymous function scope
			err = e
			return
		}
		// a non-error value was recovered
		// create a new error, 'ErrNonErrCaught', with
		// information about the recovered value
		error := fmt.Sprintf("non-error panic type %T %s", p, p)
		err = ErrNonErrCaught{Issue: error, OccurredAt: time.Now()}
		// err = fmt.Errorf("non-error panic type %[1]T %[1]s", p)
	}()

	switch input {
	case 0:
		// input was 0, return no error (nil)
		return nil
	case 1:
		// input was 1, panic with the string "one"
		panic("one")
	case 2:
		err = fmt.Errorf("Error, 2 is not a valid input")
		return err
	}

	// no case was matched
	return nil
}

// Custom Errors
type ErrNonErrCaught struct {
	Issue      string
	OccurredAt time.Time
}

func (e ErrNonErrCaught) Error() string {
	return fmt.Sprintf("[%s] - %s", e.OccurredAt, e.Issue)
}

type User struct {
	name string
}

func (u User) String() string {
	return u.name
}

type Admin struct {
	*User
	Perms map[string]bool
}
