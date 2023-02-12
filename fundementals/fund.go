package fundementals

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	// 	When imported, the image/png9 package, which contains the init() statement in
	// Listing 5.44, is called, and the package registers itself with the image package as an image
	// format.
	_ "image/jpeg"
)

func Run() {
	// chaptertwo()
	// chapterthree()
	// chapterfour()
	// chapterfive()
	chapterSix()
	chapterFive()
}

func chapterTwo() {
	// wraparound
	var maxUint8 uint8 = 11
	maxUint8 = maxUint8 * 25
	fmt.Println("value:", maxUint8)

	// Interpreted string literals are character sequences between double quotes, as in "bar".
	a := "Say \"hello\"\n\t\tto Go!\nHi!"
	fmt.Println(a)
	// Raw String Literals with Escape Characters
	b := `Say "hello"\n\t\tto Go!\n\n\nHi!`
	fmt.Println(b)

	// Runes
	// A rune is an alias for int32 and is used to represent individual characters. A rune can be
	// made up of 1 to 3 int32 values.
	c := 'A'
	fmt.Printf("%v (%T)\n", c, c)

	// Properly Iterating over Each Character in a UTF-8 String
	d := "Hello, 世界"
	for i, c := range d {
		fmt.Printf("%d: %s (%[2]T)\n", i, string(c))
	}

	// Zero Values in Go
	func() {
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("var i %T = %v\n", i, i)
		fmt.Printf("var f %T = %f\n", f, f)
		fmt.Printf("var b %T = %v\n", b, b)
		fmt.Printf("var s %T = %q\n", s, s)
	}()

	func() {
		values := func() (int, float64, bool, string) {
			return 42, 3.14, true, "hello world"
		}
		i, f, b, s := values()
		fmt.Printf("var i %T = %v\n", i, i)
		fmt.Printf("var f %T = %f\n", f, f)
		fmt.Printf("var b %T = %v\n", b, b)
		fmt.Printf("var s %T = %q\n", s, s)
	}()

	// path
	func() {
		name := "file.txt"
		ext := path.Ext(name)
		fmt.Println("Extension:", ext)

		fp := "/home/dir"
		fp = path.Join(fp, name)
		fmt.Println("Path:", fp)
	}()

	// formatting prints
	func() {
		s := "Hello, World!"
		// use '%s' to print a string
		fmt.Printf("%s\n", s)
		// use '%q' to print a string
		fmt.Printf("%q\n", s)
		d := 123
		// use '%5d' to print an integer
		// padded on the left with spaces
		// to a minimum of 5 characters wide
		fmt.Printf("Padded: '%5d'\n", d)
		// use '%5d' to print an integer
		// padded on the left with zeros
		// to a minimum of 5 characters wide
		fmt.Printf("Padded: '%05d'\n", d)
		// a number larger than the padding
		// is printed as is
		fmt.Printf("Padded: '%5d'\n", 1234567890)

		// use '%f' to print a float
		fmt.Printf("%f\n", 1.2345)
		// use '%.2f' to print a float
		// with 2 decimal places
		fmt.Printf("%.2f\n", 1.2345)

		type user struct {
			name string
			age  int
		}

		u := user{
			name: "Kurt",
			age:  27,
		}
		// use '%+v' to print an extended
		// representation of a value, if possible
		fmt.Printf("%+v\n", u)
		// use '%#v' to print the
		// Go-syntax representation of a value
		fmt.Printf("%#v\n", u)

		a, _ := strconv.Atoi("42")
		fmt.Printf("%[1]v [%[1]T]\n", a)

		b, _ := strconv.ParseFloat("42.222", 64)
		fmt.Printf("%[1]v [%[1]T]\n", b)
	}()
}

func chapterThree() {
	// Array
	func() {
		names := [4]string{"Kurt", "Janis", "Jimi", "Amy"}
		fmt.Println(names)
	}()

	// Slice
	func() {
		names := []string{"Kurt", "Janis", "Jimi", "Amy"}
		fmt.Println(names)
	}()

	// Append
	// Appending Two Slices Using the Variadic Operator
	func() {
		// create a slice of strings
		var names []string
		// append a name to the slice
		names = append(names, "Kris")
		fmt.Println(names)
		// create another slice of strings
		more := []string{"Janis", "Jimi"}
		// loop through the additional names
		names = append(names, more...)
		fmt.Println(names)
	}()

	// Printing Slice Capacity of One Million Iterations
	func() {
		var sl []int
		hat := cap(sl)
		for i := 0; i < 1_000_000; i++ {
			sl = append(sl, i)
			c := cap(sl)
			if c != hat {
				fmt.Println(hat, c)
			}
			hat = c
		}
	}()

	// Obtaining Subsets of a Slice
	func() {
		letters := []string{"a", "b", "c", "d", "e", "f", "g"}
		fmt.Println(letters) // [a b c d e f g]
		// Get 3 elements starting with the third element
		fmt.Println(letters[2:5]) // [c d e]
		// functionally equivalent
		// fmt.Println(letters[4:len(letters)]) // [e f g]
		fmt.Println(letters[4:]) // [e f g]
		// functionally equivalent
		fmt.Println(letters[0:4])             // [a b c d]
		fmt.Println(letters[:4])              // [a b c d]
		fmt.Println(letters[:1])              // [a]
		fmt.Println(letters[len(letters)-1:]) // [g]
	}()

	// Copy slices
	func() {
		names := []string{"Kurt", "Janis", "Jimi", "Amy"}
		// print the names slice
		fmt.Println(names)
		// make a new slice with
		// the correct length and
		// capacity to hold the subset
		subset := make([]string, 3)
		// copy the first three elements
		// of the names slice into the
		// subset slice
		copy(subset, names[:3])
		// print out the subset slice
		fmt.Println(subset)
		// loop over the subset slice
		for i, g := range subset {
			// uppercase each string
			// and replace the value in
			// the subset slice
			subset[i] = strings.ToUpper(g)
		}
		// print out the subset slice, again
		fmt.Println(subset)
		// print out the original names slice
		fmt.Println(names)

		slicesOnly := func(names []string) {
			for _, name := range names {
				fmt.Println(name)
			}
		}
		// convert to slice of strings
		// using the 'array[:]' syntax
		slicesOnly(names[:])
	}()

	var i int
	// create an infinite loop
	for {
		// increment the index
		i++
		if i == 3 {
			// go to the start of the loop
			continue
		}
		if i == 10 {
			// stops the loop
			break
		}
		fmt.Println(i)
	}
	fmt.Println("finished")
}

func chapterFour() {
	// Maps Have an “Unlimited” Capacity
	func() {
		users := map[string]string{}
		fmt.Println("Map length:", len(users))
		users["Kurt"] = "kurt@example.com"
		users["Janis"] = "janis@example.com"
		users["Jimi"] = "jimi@example.com"
		users["Amy"] = "Amy@example.com"
		fmt.Println("Map length:", len(users))
	}()

	func() {
		// var users map[string]string
		users := make(map[string]string)
		users["kurt@example.com"] = "Kurt"
		users["janis@example.com"] = "Janis"
		users["jimi@example.com"] = "Jimi"
		// delete the "Unknown" entry
		_, ok := users["Unknown"]
		if !ok {
			fmt.Printf("Key not found: %q\n", "Unknown")
			// os.Exit(1)
		} else {
			delete(users, "Unknown")
		}
		// print the map
		fmt.Println(users)

		// delete the "Kurt" entry
		delete(users, "kurt@example.com")
		// print the map
		fmt.Println(users)
	}()

	// Counting workds
	func() {
		counts := map[string]int{}
		sentence := "The quick brown fox jumps over the lazy dog"
		words := strings.Fields(strings.ToLower(sentence))
		for _, w := range words {
			// if the word is already in the map, increment it
			// otherwise, set it to 1 and add it to the map
			counts[w]++
		}
		fmt.Println(counts)
	}()

	// Sorting Keys and Retrieving Values from a Map
	func() {
		// create a map of months
		months := map[int]string{
			1:  "January",
			2:  "February",
			3:  "March",
			4:  "April",
			5:  "May",
			6:  "June",
			7:  "July",
			8:  "August",
			9:  "September",
			10: "October",
			11: "November",
			12: "December",
		}
		// create a slice to hold the keys
		// set its length to 0 to start with
		// and its capacity to the length
		// of the map
		keys := make([]int, 0, len(months))
		// loop through the map
		for k := range months {
			// append the key to the slice
			keys = append(keys, k)
		}

		// sort the keys
		sort.Ints(keys)
		// loop through the keys
		// and print the key/value pairs
		for _, k := range keys {
			fmt.Printf("%02d: %s\n", k, months[k])
		}
	}()

	// Scoping Variables to an if Statement
	func() {
		users := map[string]int{
			"Kurt":  27,
			"Janis": 15,
			"Jimi":  40,
		}
		name := "Amy"
		if age, ok := users[name]; ok {
			fmt.Printf("%s is %d years old\n", name, age)
			return
		}
		fmt.Printf("Couldn't find %s in the users map\n", name)
	}()

	// Using fallthrough with a switch Statement
	func() {
		recommendActivity := func(temp int) {
			fmt.Printf("It is %d degrees out. You could", temp)
			switch {
			case temp <= 32:
				fmt.Print(" go ice skating,")
				fallthrough
			case temp >= 45 && temp < 90:
				fmt.Print(" go jogging,")
				fallthrough
			case temp >= 80:
				fmt.Print(" go swimming,")
				fallthrough
			default:
				fmt.Print(" or just stay home.\n")
			}
		}
		recommendActivity(19)
		recommendActivity(45)
		recommendActivity(90)
	}()
}

func decode(reader io.Reader) (image.Rectangle, error) {
	// decode the image reader
	m, _, err := image.Decode(reader)
	if err != nil {
		// return the error
		return image.Rectangle{}, err
	}
	return m.Bounds(), nil
}

func chapterFive() {
	// Deferred Function Calls Are Executed in LIFO Order
	reader, err := os.Open("data/pix.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}

	func() {
		defer fmt.Println("one")
		defer fmt.Println("two")
		defer fmt.Println("three")

		defer func() {
			fmt.Println("closing")
			// src.Close()
		}()

		// Deferred Calls Are Executed Even if Another Deferred Call Panics
		// defer fmt.Println("one")
		// defer panic("two")
		// defer fmt.Println("three")

		// // Deferred Calls Are Not Executed if the Code Exits
		// defer fmt.Println("one")
		// os.Exit(1)
		// defer fmt.Println("three")

		// // Deferred Calls Are Not Executed if the Code Logs a Fatal Message
		// defer fmt.Println("one")
		// log.Fatal("boom")
		// defer fmt.Println("three")
	}()

	// capture the current time
	now := time.Now()
	// use an anonymous function
	// to scope variables to be
	// used in the defer
	fmt.Printf("0 duration: %s\n", time.Since(now))
	defer fmt.Printf("5 duration: %s\n", time.Since(now))
	defer func(now time.Time) {
		fmt.Printf("4 duration: %s\n", time.Since(now))
	}(now)
	defer func() {
		fmt.Printf("3 duration: %s\n", time.Since(now))
	}()
	defer func(now time.Time) {
		fmt.Printf("2 duration: %s\n", time.Since(now))
	}(now)
	defer func() {
		fmt.Printf("1 duration: %s\n", time.Since(now))
	}()
	fmt.Println("sleeping for 50ms...")
	// sleep for 50ms
	time.Sleep(50 * time.Millisecond)
}

// Using the json Struct Tag to Control Encoding Output
type user struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func (u user) String() string {
	return u.Name + " Email:" + u.Email
}

// Pointer Receiver
func (u *user) Titleize() {
	u.Name = strings.ToUpper(u.Name)
}

func chapterSix() {
	func() {
		type myInt int
		type myString string
		type myMap map[string]string

		// declare a variable of type MyInt
		i := myInt(1)
		// declare a variable of type MyString
		s := myString("foo")
		// declare a variable of type MyMap
		m := myMap{"foo": "bar"}
		// print the type and value of i
		fmt.Printf("%[1]T:\t%[1]v\n", i)
		// print the type and value of s
		fmt.Printf("%[1]T:\t%[1]v\n", s)
		// print the type and value of m
		fmt.Printf("%[1]T:\t%[1]v\n", m)

		u := user{}
		n, e := u.Name, u.Email
		fmt.Printf("%+v%+v%+v\n", u, n, e)

		u = user{Name: "Ben", Email: "brgeyer49@gmail.com", Password: "topsecret"}
		fmt.Println(u.String())
		u.Titleize()
		fmt.Println(u.String())

		enc := json.NewEncoder(os.Stdout)
		// encode the user
		if err := enc.Encode(u); err != nil {
			// handle an error if one occurs
			log.Fatal(err)
		}
	}()

	func() {
		// create a pointer to a string
		s := new(string)
		s2 := "hello world"
		// dereference the pointer
		// and assign a value to it
		*s = "hello"
		// create a pointer to an int
		i := new(int)
		// dereference the pointer
		// and assign a value to it
		*i = 42
		// create a pointer to a user
		u1 := new(user)
		// functionally equivalent and idiomatic
		u2 := &user{Email: "jammin"}

		// mutate string pointer 1
		// cap := func(s *string) {
		// 	c := strings.ToUpper(*s)
		// 	*s = c
		// }

		// mutate string pointer 2
		// cap := func(s *string) {
		// 	c := new(string)
		// 	*c = strings.ToUpper(*s)
		// 	*s = *c
		// }

		// mutate string pointer 3
		cap := func(s *string) {
			// important to make the pointer is not nil
			if s != nil {
				*s = strings.ToUpper(*s)
			}
		}

		fmt.Println("Pointer mutation strings")
		fmt.Println(s2)
		cap(&s2)
		fmt.Println(s2)
		cap(s)
		cap(&u2.Email)

		fmt.Println(*s)
		fmt.Printf("s: %v, *s: %q\n", s, *s)
		fmt.Printf("i: %v, *i: %d\n", i, *i)
		fmt.Printf("u1: %+v, *f: %+v\n", u1, *u1)
		fmt.Printf("u2: %+v, *f1: %+v\n", u2, *u2)
	}()
}
