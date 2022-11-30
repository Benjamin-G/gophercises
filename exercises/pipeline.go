package exercises

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
)

func makeChannel(nums *[]int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range *nums {
			out <- v
			// fmt.Println("makeChannel :", v)
			// this channel will run synchronously with the pipeline as it is blocked until it is read
		}
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
		// close(out)
	}()

	return out
}

func Pipeline() {
	// https://www.youtube.com/watch?v=qyM8Pi1KiiM&t=59s
	nums := []int{2, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	stage1 := makeChannel(&nums)

	stage2 := makeDouble(stage1)

	stage3 := makeNumStruct(stage2)

	for v := range stage3 {
		fmt.Printf("Value: %v\n", v)
	}

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
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data2 := []byte("golang")
	go printData(&wg, data2[:3])
	go printData(&wg, data2[3:])

	wg.Wait()
}
