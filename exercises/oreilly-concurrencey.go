package exercises

import (
	"fmt"
)

// Let’s first define Little’s Law algebraicly. It is commonly expressed as: L=λW, where:

// L = the average number of units in the system.

// λ = the average arrival rate of units.

// W = the average time a unit spends in the system.

// This equation only applies to so-called stable systems. In a pipeline, a stable system is one in which the rate that work enters the pipeline, or ingress, is equal to the rate in which it exits the system, or egress. If the rate of ingress exceeds the rate of egress, your system is unstable and has entered a death-spiral. If the rate of ingress is less than the rate of egress, you still have an unstable system, but all that’s happening is that your resources aren’t being utilized completely. Not the worst situation in the world, but maybe you care about this if the underutilization is found on a vast scale (e.g., clusters or data centers).

func OriellyCurrent() {

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

	orDone := func(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_ <-chan interface{}, _ <-chan interface{}, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		out3 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			defer close(out3)
			for val := range orDone(done, in) {
				var out1, out2, out3 = out1, out2, out3
				for i := 0; i < 3; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					case out3 <- val:
						out3 = nil
					}
				}
			}
		}()
		return out1, out2, out3
	}

	done := make(chan interface{})
	defer close(done)

	// Tee is returning two channels, each will send a value two each channel
	out1, out2, out3 := tee(done, take(done, repeat(done, 1, 2), 4))

	// Utilizing this pattern, it’s easy to continue using channels as the join points of your system.
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v, out3:%v\n", val1, <-out2, <-out3)
	}

	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	// This is pretty straightforward code. Now we can use bridge to help present a single-channel facade over a channel of channels. Here’s an example that creates a series of 10 channels, each with one element written to them, and passes these channels into the bridge function:
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(done, genVals()) {
		fmt.Printf("%v ", v)
	}

}
