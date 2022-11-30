package exercises

import "fmt"

func makeChannel(nums *[]int) chan int {
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

func makeDouble(in <-chan int) chan int {
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

func makeNumStruct(in <-chan int) chan NumStruct {
	out := make(chan NumStruct)

	go func() {
		for v := range in {
			out <- NumStruct{id: v / 2, double: v}
		}
		close(out)
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

	// data := make([]int, 4)

	// loopData := func(handleData chan<- int) {
	// 	defer close(handleData)
	// 	for i := range data {
	// 		handleData <- data[i]
	// 	}
	// }

	// handleData := make(chan int)
	// go loopData(handleData)

	// for num := range handleData {
	// 	fmt.Println(num)
	// }
}
