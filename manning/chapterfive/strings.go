package chapterfive

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func ChpFiveRunner() {
	substring()
	runes()
	trim()
}

func substring() {
	printLength := func(s string) {
		fmt.Printf("%s length(bytes): %d \n", s, len(s))
	}
	s1 := "Hello, World!"
	printLength(s1)
	//This example creates a string from the first five
	//bytes, not the first five runes.
	s2 := s1[:5]
	printLength(s2)

	s1 = "Hêllo, World!"
	printLength(s1)
	s2 = string([]rune(s1)[:5])
	s3 := s1[:5]
	s4 := strings.Clone(string([]rune(s1)[:5]))
	printLength(s2)
	printLength(s3)
	printLength(s4)
}

func runes() {
	s := "hello"
	fmt.Println(len(s))

	s = "汉"
	fmt.Println(len(s))
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s)

	s = "hêllo"
	fmt.Printf("%s\n", s)
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))

	fmt.Printf("%s\n", s)
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	runes := []rune(s)
	fmt.Printf("%v\n", runes)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4]))
}

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}

func trim() {

	fmt.Println(strings.TrimRight("123oxo", "xo"))
	fmt.Println(strings.TrimRight("123oxo", "o"))   // 123ox
	fmt.Println(strings.TrimLeft("oxo123", "ox"))   // 123
	fmt.Println(strings.TrimPrefix("oxo123", "ox")) /// o123
}

func concat1(values []string) string {
	//But with this implementation, we forget one
	//of the core characteristics of a string: its immutability. Therefore, each iteration
	//doesn’t update s; it reallocates a new string in memory, which significantly impacts the
	//performance of this function.
	s := ""
	for _, value := range values {
		s += value
	}

	return s
}

func concat2(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func concat3(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}
	//Indeed, if we just have to concatenate
	//a few strings (such as a name and a surname), using strings.Builder is not recommended
	//as doing so will make the code a bit less readable than using the +=
	//	operator or fmt.Sprintf.

	sb := strings.Builder{}
	sb.Grow(total)
	//Note that we’re not interested in the number of runes
	//but the number of bytes, so we use the len function. Then we call Grow to guarantee
	//space for total bytes before iterating over the strings.
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func getBytes1(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []byte(sanitize1(string(b))), nil
}

func sanitize1(s string) string {
	return strings.TrimSpace(s)
}

func getBytes2(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize2(b), nil
}

func sanitize2(b []byte) []byte {
	return bytes.TrimSpace(b)
}
