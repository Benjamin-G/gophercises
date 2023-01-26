package fundementals

import (
	"fmt"
	"time"
)

// Messages Are Pulled off a Channel by the First Goroutine to Read from It
func ChannelsMain() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("[Fatal] invalid matcher. panic occurred: %v\n", e)
		}
	}()

	const N = 5
	// make a new channel
	// of type int
	ch := make(chan int)
	for i := 0; i < N; i++ {
		// create a goroutine to listen to the channel
		go func(i int) {
			// listen for new messages
			for m := range ch {
				fmt.Printf("routine %d received %d\n", i, m)
			}
		}(i)
	}
	// print messages to the channel
	for i := 0; i < N*2; i++ {
		ch <- i
	}
	// close the channel
	// this will break the 'range'
	// statement in the goroutine
	close(ch)
	// wait for the goroutines to finish
	// time.Sleep(50 * time.Millisecond)

	// create a channel to signal listeners to exit
	quit := make(chan struct{})
	// create 5 listeners
	ch2 := make(chan int)
	go listener(ch2, quit)
	// for i := 0; i < 5; i++ {
	// 	// launch listener in a goroutine
	// 	go listener(ch2, quit)
	// }

	time.Sleep(10 * time.Millisecond)
	for i := 0; i < 500; i++ {
		// launch listener in a goroutine
		ch2 <- i
	}
	defer close(ch2)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("closing the quit channel")
	// close the channel to signal listeners to exit
	close(quit)

	// reading from a closed channel
	// make a buffered channel of ints
	// that can hold 5 values before blocking
	ch3 := make(chan int, 5)
	// write messages to the channel
	for i := 0; i < 5; i++ {
		ch3 <- i
	}
	// close the channel
	close(ch3)
	// we can continue to read messages
	// from the closed channel until it is empty.
	// when it is empty the for loop will exit.
	for i := range ch3 {
		fmt.Println(i)
	}

	// allow the listeners to exit
	time.Sleep(10 * time.Millisecond)
}

func listener(ch <-chan int, quit <-chan struct{}) {
	// infinite loop to keep listening
	// for messages on the channel
	for {
		// store the message from the channel to variable i
		// capture if the channel is closed or not to variable ok
		select {
		case i, ok := <-ch:
			// if the channel is closed, return from the function
			if !ok {
				fmt.Println("closed channel")
				return
			}
			// print the message
			fmt.Printf("read %d from channel\n", i)
		case <-quit:
			fmt.Println("closed channel from quit")
			return
		}
	}
}

type Newspaper struct {
	headlines chan string
	quit      chan struct{}
}

// TopHeadlines returns a read-only channel of strings
// that represent the top headlines of the newspaper.
// This channel is consumed by newspaper readers.
func (n Newspaper) TopHeadlines() <-chan string {
	return n.headlines
}

// ReportStory returns a write-only channel of strings
// that a reporter can use to report a story.
func (n Newspaper) ReportStory() chan<- string {
	return n.headlines
}
