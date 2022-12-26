package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
)

// go test -bench=. test/buffering_test.go

// If batching requests in a stage saves time.

// If delays in a stage produce a feedback loop into the system.

// BenchmarkUnbufferedWrite-8         78495             13899 ns/op
// BenchmarkBufferedWrite-8          591888              2284 ns/op
// go test -bench=. tests/buffering_test.go -count 5
// On the right side of the function name, you have two values, 14588 and 82798 ns/op. The former indicates the total number of times the loop was executed, while the latter is the average amount of time each iteration took to complete, expressed in nanoseconds per operation.

// Usually anytime performing an operation requires an overhead, chunking may increase system performance. Some examples of this are opening database transactions, calculating message checksums, and allocating contiguous space.

// By introducing a queue at the entrance to the pipeline, you can break the feedback loop at the cost of creating lag for requests.

// If the caller does time out, you need to be sure you support some kind of check for readiness when dequeuing.

// You may be tempted to add queuing elsewhere—e.g., after a computationally expensive stage—but avoid that temptation!

// Again, our pipeline has three stages, so we’ll decrement L by 3. We set λ to 100,000 r/s, and find that if we want to field that many requests, our queue should have a capacity of 7. Remember that as you increase the queue size, it takes your work longer to make it through the system! You’re effectively trading system utilization for lag.

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferredFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferredFile))
}

func tmpFileOrFatal() *os.File {
	file, err := os.CreateTemp("", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
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

	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}
