package fundementals

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"
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
		debug.PrintStack()
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

type ErrorA struct {
	err error
}

func (e ErrorA) Error() string {
	return fmt.Sprintf("[ErrorA] %s", e.err)
}

func (e ErrorA) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}

	return e.err
}

func (e ErrorA) As(target any) bool {
	ex, ok := target.(*ErrorA)
	if !ok {
		// if the target is not an ErrorA,
		// pass the underlying error up the chain
		// by calling errors.As with the underlying error
		// and the target error
		return errors.As(e.err, target)
	}
	// set the target to the current error
	(*ex) = e
	return true
}

func (e ErrorA) Is(target error) bool {
	if _, ok := target.(ErrorA); ok {
		// return true if target is ErrorA
		return true
	}
	// if not, pass the underlying error up the chain
	// by calling errors.Is with the underlying error
	// and the target error
	return errors.Is(e.err, target)
}

type ErrorB struct {
	err error
}

func (e ErrorB) Error() string {
	return fmt.Sprintf("[ErrorB] %s", e.err)
}

func (e ErrorB) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}

	return e.err
}

func (e ErrorB) As(target any) bool {
	ex, ok := target.(*ErrorB)
	if !ok {
		// if the target is not an ErrorA,
		// pass the underlying error up the chain
		// by calling errors.As with the underlying error
		// and the target error
		return errors.As(e.err, target)
	}
	// set the target to the current error
	(*ex) = e
	return true
}

func (e ErrorB) Is(target error) bool {
	if _, ok := target.(ErrorB); ok {
		// return true if target is ErrorA
		return true
	}
	// if not, pass the underlying error up the chain
	// by calling errors.Is with the underlying error
	// and the target error
	return errors.Is(e.err, target)
}

type ErrorC struct {
	err error
}

func (e ErrorC) Error() string {
	return fmt.Sprintf("[ErrorC] %s", e.err)
}

func (e ErrorC) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}

	return e.err
}

func (e ErrorC) As(target any) bool {
	ex, ok := target.(*ErrorC)
	if !ok {
		// if the target is not an ErrorA,
		// pass the underlying error up the chain
		// by calling errors.As with the underlying error
		// and the target error
		return errors.As(e.err, target)
	}
	// set the target to the current error
	(*ex) = e
	return true
}

func (e ErrorC) Is(target error) bool {
	if _, ok := target.(ErrorC); ok {
		// return true if target is ErrorA
		return true
	}
	// if not, pass the underlying error up the chain
	// by calling errors.Is with the underlying error
	// and the target error
	return errors.Is(e.err, target)
}

// Like the errors.Unwrap function, errors.As also has a documented, but unpublished,
// interface, Listing 9.64, that can be implemented on custom errors.
type AsError interface {
	As(target any) bool
}

type Unwrapper interface {
	Unwrap() error
}

type IsError interface {
	Is(target error) bool
}

// Wrapper wraps an error with a bunch of
// other errors.
// ex. Wrapper(original) #=> ErrorC -> ErrorB -> ErrorA -> original
func Wrapper(original error) error {
	original = ErrorA{original}
	original = ErrorB{original}
	original = ErrorC{original}
	return original
}

// WrapperLong wraps an error with a bunch of
// other errors.
// ex. WrapperLong(original) #=> ErrorC -> ErrorB -> ErrorA -> original
func WrapperLong(original error) error {
	return ErrorC{
		err: ErrorB{
			err: ErrorA{
				err: original,
			},
		},
	}
}
