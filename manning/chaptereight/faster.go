package chaptereight

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"sync/atomic"
)

func sequentialMergesort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle])
	sequentialMergesort(s[middle:])
	merge(s, middle)
}

func parallelMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)
}

const max = 2048

// This threshold will represent
// how many elements a half should contain in order to be handled in a parallel manner.
// If the number of elements in the half is fewer than this value, we will handle it sequentially.
// Here’s a new version:

func parallelMergesortV2(s []int) {
	if len(s) <= 1 {
		return
	}

	if len(s) <= max {
		sequentialMergesort(s)
	} else {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func giveTwo() int64 {
	var i int64

	ch := make(chan int64)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch

	return i
}

// This is so weird. I don't understand why this does not work.
func giveTwoV2() int64 {
	var i atomic.Int64

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		i.Add(1)
		i.Add(1)
	}()

	func() {
		i := 0
		ch := make(chan struct{})

		go func() {
			<-ch
			fmt.Println(i)
		}()
		i++
		ch <- struct{}{}
	}()

	wg.Wait()

	return i.Load()
}

// workload type
func dumbReaderRun() {
	res1, _ := read1(&dummyReader{})
	fmt.Println(res1)

	res2, _ := read2(&dummyReader{})
	fmt.Println(res2)
}

func read1(r io.Reader) (int, error) {
	count := 0

	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)

		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}

		count += task(b)
	}

	return count, nil
}

func read2(r io.Reader) (int, error) {
	var count int64

	wg := sync.WaitGroup{}

	n := runtime.GOMAXPROCS(0)
	fmt.Printf("GOMAXPROCS: %d\n", n)
	//n := 10

	ch := make(chan []byte, n)
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)

		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}
		ch <- b
	}

	close(ch)
	wg.Wait()

	return int(count), nil
}

func task(b []byte) int {
	return len(b)
}

type dummyReader struct {
	i int
}

func (c *dummyReader) Read(p []byte) (n int, err error) {
	if c.i == 3 {
		return 0, io.EOF
	}

	copy(p, []byte{0, 1, 2})
	c.i++

	return 3, nil
}
