package fundementals

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func Test_ThumbnailGenerator(t *testing.T) {
	t.Parallel()
	// image that we need thumbnails for
	const image = "foo.png"
	var wg sync.WaitGroup
	// start 5 goroutines to generate thumbnails
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// start a new goroutine for each thumbnail
		go generateThumbnail(&wg, image, i+1)
	}
	fmt.Println("Waiting for thumbnails to be generated")
	// wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Finished generate all thumbnails")
}

func Test_ThumbnailGeneratorV2(t *testing.T) {
	t.Parallel()
	// image that we need thumbnails for
	const image = "foo.png"
	// create a new error group
	var wg errgroup.Group
	// start 5 goroutines to generate thumbnails
	for i := 0; i < 5; i++ {
		// capture the i to the current scope
		i := i
		// start a new goroutine for each thumbnail
		wg.Go(func() error {
			// return the result of generateThumbnail
			return generateThumbnailV2(image, i)
		})
	}
	fmt.Println("Waiting for thumbnails to be generated")
	// wait for all goroutines to finish
	err := wg.Wait()
	// check for any errors
	if err != nil {
		fmt.Println(err)
		// t.Fatal(err)
	}
	fmt.Println("Finished generate all thumbnails")
}

func Test_ErrorGroup_Multiple_Errors(t *testing.T) {
	t.Parallel()
	var wg errgroup.Group
	for i := 0; i < 10; i++ {
		i := i + 1
		wg.Go(func() error {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			fmt.Printf("about to error from %d\n", i)
			return fmt.Errorf("error %d", i)
		})
	}
	err := wg.Wait()
	if err != nil {
		fmt.Println(err)
		// t.Fatal(err)
	}
}

func Test_WaitGroup_Add_Positive(t *testing.T) {
	t.Parallel()
	var completed bool
	// create a new waitgroup (count: 0)
	var wg sync.WaitGroup
	// add one to the waitgroup (count: 1)
	wg.Add(1)
	// launch a goroutine to call the Done() method
	go func(wg *sync.WaitGroup) {
		// sleep for a bit
		time.Sleep(time.Millisecond * 10)
		fmt.Println("done with waitgroup")
		completed = true
		// call the Done() method to decrement
		// the waitgroup counter (count: 0)
		wg.Done()
	}(&wg)
	fmt.Println("waiting for waitgroup to unblock")
	// wait for the waitgroup to unblock (count: 1)
	wg.Wait()
	// (count: 0)
	fmt.Println("waitgroup is unblocked")
	if !completed {
		t.Fatal("waitgroup is not completed")
	}
}

func Test_WaitGroup_Add_Zero(t *testing.T) {
	t.Parallel()
	// create a new waitgroup (count: 0)
	var wg sync.WaitGroup
	// add 0 to the waitgroup (count: 0)
	wg.Add(0)
	// (count: 0)
	fmt.Println("waiting for waitgroup to unblock")
	// wait for the waitgroup to unblock (count: 0)
	// will not block since the counter is alrady 0
	wg.Wait()
	// (count: 0)
	fmt.Println("waitgroup is unblocked")
}

func Test_WaitGroup_Done(t *testing.T) {
	t.Parallel()
	const N = 5
	// create a new waitgroup (count: 0)
	var wg sync.WaitGroup
	// add 5 to the waitgroup (count: 5)
	wg.Add(N)
	for i := 0; i < N; i++ {
		// launch a goroutine that will call the
		// waitgroup's Done method when it finishes
		go func(i int) {
			// sleep briefly
			time.Sleep(time.Millisecond * time.Duration(i))
			fmt.Println("decrementing waiting by 1")
			// call the waitgroup's Done method
			// (count: count - 1)
			wg.Done()
		}(i + 1)
	}
	fmt.Println("waiting for waitgroup to unblock")
	wg.Wait()
	fmt.Println("waitgroup is unblocked")
}

func Test_ErrorGroup_Context(t *testing.T) {
	t.Parallel()
	// create a new error group
	// and a context that will be canceled
	// when the group is done
	wg, ctx := errgroup.WithContext(context.Background())
	// create a quit channel for the goroutine
	// waiting for the context to be canceled
	// can close to signal the goroutine has finished
	quit := make(chan struct{})
	// launch a goroutine that will
	// wait for the errgroup context to finish
	go func() {
		fmt.Println("waiting for context to cancel")
		// wait for the context to be canceled
		<-ctx.Done()
		fmt.Println("context canceled")
		// close the quit channel so the test
		// will finish
		close(quit)
	}()
	// add a task to the errgroup
	wg.Go(func() error {
		time.Sleep(time.Millisecond * 5)
		return nil
	})
	// wait for the errgroup to finish
	err := wg.Wait()
	if err != nil {
		t.Fatal(err)
	}
	// wait for the context goroutine to finish
	<-quit
}

func Test_Mutex(t *testing.T) {
	t.Parallel()
	var mu sync.RWMutex
	// create a new cancellable context
	// to stop the test when the goroutines
	// are finished
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Millisecond)
	defer cancel()
	// create a map to be used
	// as a shared resource
	data := map[int]bool{}
	// launch a goroutine to
	// write data in the map
	go func() {
		// lock the mutex
		mu.Lock()
		for i := 0; i < 10; i++ {
			// loop putting data in the map
			data[i] = true
		}
		// unlock the mutex
		mu.Unlock()
		// cancel the context
		cancel()
	}()
	// launch a goroutine to
	// read data from the map
	go func() {
		// lock the mutex
		mu.RLock()
		// loop through the map
		// and print the keys/values
		for k, v := range data {
			fmt.Printf("%d: %v\n", k, v)
		}
		// unlock the mutex
		mu.RUnlock()
	}()
	// wait for the context to be canceled
	<-ctx.Done()
	if len(data) != 10 {
		t.Fatalf("expected 10 items in the map, got %d", len(data))
	}
}

func Test_Once(t *testing.T) {
	t.Parallel()
	b := &Builder{}
	for i := 0; i < 5; i++ {
		err := b.Build()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("builder built")
		if !b.Built {
			t.Fatal("expected builder to be built")
		}
	}
}

func Test_Closing_Channels(t *testing.T) {
	t.Parallel()
	func() {
		// defer a function to catch the panic
		defer func() {
			// recover the panic
			if r := recover(); r != nil {
				// mark the test as a failure
				t.Fatal(r)
			}
		}()
		m := &Manager{
			quit: make(chan struct{}),
		}
		// close the manager's quit channel
		m.Quit()
		// try to close the manager's quit channel again
		// this will panic
		m.Quit()
	}()
}
