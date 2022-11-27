package exercises

import (
	"fmt"
	"sync"
	"time"
)

// https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
// Using goroutines is a very quick way to turn what would be a sequential program into a concurrent program without having to worry about things like creating threads or thread-pools.

// Goroutines are incredibly lightweight “threads” managed by the go runtime.

// Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes 1Mb.

// Creating a thousand goroutines would typically require one or two OS threads at most, whereas if we were to do the same thing in java it would require 1,000 full threads each taking a minimum of 1Mb of Heap space.

// a very simple function that we'll
// make asynchronous later on
func compute(value int, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i < value; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i)
	}
	wg.Done()
}

func newGenericFunc[age int64 | float64, anon any](myAge age, sec anon) {
	fmt.Println(myAge, sec)
}

func logg[input any](i input) {
	fmt.Println(i)
}

type Number interface {
	int16 | int32 | int64 | float32 | float64
}

func bubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

// Interfaces with structs
type Service interface {
	SayHi()
}

type MyService struct{}

func (s MyService) SayHi() {
	fmt.Println("Hi")
}

type SecondService struct {
	id int
}

func (s SecondService) SayHi() {
	fmt.Println("Hello From the 2nd Service ", s.id)
}

func Goroutine() {
	fmt.Println("Goroutine Tutorial")

	var wg sync.WaitGroup

	// sequential execution of our compute function
	go compute(10, &wg)
	// go compute(10, &wg)

	fmt.Println("Wait for routines to finish")
	wg.Wait()
	fmt.Println("Done!")

	// we scan fmt for input and print that to our console
	// var input string
	// fmt.Scanln(&input)

	// we have to once again block until our anonymous goroutine
	// has finished or our main() function will complete without
	// printing anything
	// fmt.Scanln()
	var testAge int64 = 23
	var testAge2 float64 = 24.5
	newGenericFunc(testAge, "Hello, world!")
	newGenericFunc(testAge2, testAge)

	wg.Add(1)
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
		wg.Done()
	}()
	wg.Wait()
	list := []int32{4, 3, 1, 5}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2}
	sorted := bubbleSort(list)
	fmt.Println(sorted)

	sortedFloats := bubbleSort(list2)
	fmt.Println(sortedFloats)

	// MAPs
	// youtubeSubscribers := map[string]int{
	// 	"TutorialEdge":     2240,
	// 	"MKBHD":            6580350,
	// 	"Fun Fun Function": 171220,
	// }

	youtubeSubscribers := map[string]int{}
	youtubeSubscribers["abc"] = 2240
	// fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350
	logg(youtubeSubscribers["abc"])

	// we can define a map of string uuids to
	// the interface type 'Service'
	// interfaceMap := make(map[string]Service)

	// We use the interface to allow any struct with
	interfaceMap := map[string]Service{}

	// we can then populate our map with
	// simple ids to particular services
	interfaceMap["SERVICE-ID-1"] = MyService{}
	interfaceMap["SERVICE-ID-2"] = SecondService{id: 123}

	// Incoming HTTP Request wants service 2
	// we can use the incoming uuid to lookup the required
	// service and call it's SayHi() method
	interfaceMap["SERVICE-ID-2"].SayHi()
}
