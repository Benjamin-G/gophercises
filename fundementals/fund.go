package fundementals

import "fmt"

func Run() {
	chapterTwo()
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
		fmt.Printf("%d: %s (%T)\n", i, string(c), string(c))
	}
}
