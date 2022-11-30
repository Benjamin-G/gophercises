package exercises

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func makeChannel(nums *[]int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range *nums {
			out <- v
			// fmt.Println("makeChannel :", v)
			// this channel will run synchronously with the pipeline as it is blocked until it is read
		}
		fmt.Println("makeChannel exited.")
		close(out)
	}()

	return out
}

func makeDouble(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// this is going to read from the channel until line 14 is reached
		for v := range in {
			out <- v * v
			// fmt.Println("makeDouble :", v*v)
		}
		fmt.Println("makeDouble exited.")
		close(out)
	}()

	return out
}

type NumStruct struct {
	id     int
	double int
}

func makeNumStruct(in <-chan int) <-chan NumStruct {
	out := make(chan NumStruct)

	go func() {
		defer close(out)
		for v := range in {
			out <- NumStruct{id: v / 2, double: v}
		}
		fmt.Println("makeNumStruct exited.")
		// close(out)
	}()

	return out
}

func Pipeline() {
	var wg sync.WaitGroup
	// https://www.youtube.com/watch?v=qyM8Pi1KiiM&t=59s
	nums := []int{2, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	wg.Add(1)
	go func() {
		stage1 := makeChannel(&nums)

		stage2 := makeDouble(stage1)

		stage3 := makeNumStruct(stage2)

		for v := range stage3 {
			fmt.Printf("Value: %v\n", v)
		}
		wg.Done()
	}()

	// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i] + rand.Intn(100)
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}

	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer fmt.Println("printData exited.")
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	wg.Add(2)
	data2 := []byte("golang")
	go printData(&wg, data2[:3])
	go printData(&wg, data2[3:])

	wg.Wait()

	// Now that we know how to ensure goroutines don’t leak, we can stipulate a convention: If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
			// fmt.Println("This will not run.")
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}
