package exercises

import (
	"fmt"
	"net/http"
	"time"
)

func Oreilly() {
	fmt.Println("Or Pattern for multiple Channels")
	//line 36 requires the variable initialized, recursion
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		default:
			fmt.Println(len(channels))
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	// that one of Go’s strengths is the ability to quickly create, schedule, and run goroutines, and the language actively encourages using goroutines to model problems correctly.

	// This pattern is useful to employ at the intersection of modules in your system. At these intersections, you tend to have multiple conditions for canceling trees of goroutines through your call stack. Using the or function, you can simply combine these together and pass it down the stack. We’ll take a look at another way of doing this in “The context Package” that is also very nice, and perhaps a bit more descriptive.

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			// defer fmt.Println(after)
			defer close(c)
			time.Sleep(after)
			fmt.Println(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n\n", time.Since(start))

	fmt.Println("Error Handling")

	checkStatus := func(
		done <-chan interface{},
		urls ...string,
		// urls []string, cannot use ... in call to non-variadic checkStatuscompilerNonVariadicDotDotDot
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			// When Go eschewed the popular exception model of errors, it made a statement that error handling was important, and that as we develop our programs, we should give our error paths the same attention we give our algorithms
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost", "https://www.google.com"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
