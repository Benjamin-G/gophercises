package exercises

import (
	"fmt"
	"math/rand"
)

func Generators() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	// Empty interfaces are a bit taboo in Go, but for pipeline stages it is my opinion that it’s OK to deal in channels of interface{} so that you can use a standard library of pipeline patterns.

	done := make(chan interface{})
	defer close(done)

	// v := []int{1, 2, 3, 4}
	v := "Hello, world!"
	for num := range take(done, repeat(done, v), 1) {
		fmt.Printf("%v ", num)
	}

	rand := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFn(done, rand), 1) {
		fmt.Println(num)
	}

	var message string
	// This is pretty neat, repeat then take a certain amount from that inifinty looping interface
	for token := range toString(done, take(done, repeat(done, "I", " am", " actually a dog."), 4)) {
		message += token
	}

	fmt.Printf("message: %s...", message)

	// So let’s prove to ourselves that the performance cost of genericizing portions of our pipeline is negligible. We’ll write two benchmarking functions: one to test the generic stages, and one to test the type-specific stages:
}
