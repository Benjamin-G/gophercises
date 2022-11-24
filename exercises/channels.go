package exercises

// https://tutorialedge.net/golang/go-channels-tutorial/

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
}

// channel <- value - sends a value to a channel

// value := <- channel - receives a value from a channel

func CalculateValueV2(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value V2: {}", value)
	// the go routine will run synchronously and concurrently without waiting
	// with the waiting the concurrence of the sends are less predictable
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("1, Only Executes after another goroutine performs a receive on the channel")
	c <- value
	fmt.Println("2, Only Executes after another goroutine performs a receive on the channel")
}

func Channels() {
	fmt.Println("Go Channel Tutorial")

	// blocks within our goroutines should the channel be full.
	values := make(chan int, 4)
	// Very strang behavior in with an unbuffered channel being sent and receiving multiple times
	// The goroutine will block for the channel to receive the send from the channel
	// The order in which the variables receive is unknown
	// values := make(chan int)

	// defer close(values)

	go CalculateValueV2(values)
	go CalculateValueV2(values)

	v1 := <-values
	v2 := <-values
	// If you try to receive without a send:
	// fatal error: all goroutines are asleep - deadlock!
	// <-values
	// <-values
	// <-values
	fmt.Println(v1, v2)
	// If a buffered channel, if given the time it will finish the goroutine
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(<-values, <-values)

}
