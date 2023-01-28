package fundementals

import (
	"context"
	"fmt"
	"os"
	"time"
)

// const (ID )
type IDtype string

var ID IDtype

// ID = "ID"
// CtxKeyA is used to wrap keys
// associated with a A request
// CtxKeyA("request_id")
// CtxKeyA("user_id")
type CtxKeyA string

// CtxKeyB is used to wrap keys
// associated with a B request
// CtxKeyB("request_id")
// CtxKeyB("user_id")
type CtxKeyB string

func ContextMain() {
	ctx := context.Background()
	// print the current value
	// of the context
	fmt.Printf("%v\n", ctx)
	// print Go-syntax representation of the value
	fmt.Printf("\t%#v\n", ctx)
	// print the value of the Done channel
	// does not block because we are not
	// trying to read/write to the channel
	fmt.Printf("\tDone:\t%#v\n", ctx.Done())
	// print the value of the Err
	fmt.Printf("\tErr:\t%#v\n", ctx.Err())
	// print the value of "KEY"
	fmt.Printf("\tValue:\t%#v\n", ctx.Value("ID"))
	// print the deadline time
	// and true/false if there is no deadline
	deadline, ok := ctx.Deadline()
	fmt.Printf("\tDeadline:\t%s (%t)\n", deadline, ok)

	// of application shutdown/cancellation.
	ctx, cancel := context.WithCancel(ctx)
	// ensure the cancel function is called
	// to shut down the monitor when the program
	// exits
	defer cancel()
	// launch a goroutine to cancel the application
	// context after a short while.
	go func() {
		time.Sleep(time.Millisecond * 50)
		// cancel the application context
		// this will shut the monitor down
		cancel()
	}()
	// create a new monitor
	mon := Monitor{}
	// start the monitor with the application context
	// this will return a context that can be listened to
	// for cancellation signaling the monitor has shut down.
	ctx = mon.Start(ctx)
	// block the application until either the context
	// is canceled or the application times out
	select {
	case <-ctx.Done(): // listen for context cancellation
		// success shutdown
		os.Exit(0)
	case <-time.After(time.Second * 2): // timeout after 2 second
		fmt.Println("timed out while trying to shut down the monitor")
		// check if there was an error from the
		// monitor's context
		if err := ctx.Err(); err != nil {
			fmt.Printf("error: %s\n", err)
		}
		// non-successful shutdown
		os.Exit(1)
	}
}

type Monitor struct {
	cancel context.CancelFunc
}

func (m *Monitor) Start(ctx context.Context) context.Context {
	// start the monitor with the given context
	go m.listen(ctx)
	// create a new context that will be canceled
	// when the monitor is shut down
	ctx, cancel := context.WithCancel(context.Background())
	// hold on to the cancellation function
	// when context that started the manager is canceled
	// this cancellation function will be called.
	m.cancel = cancel
	// return the new, cancellable, context.
	// clients can listen to this context
	// for cancellation to ensure the
	// monitor is properly shut down.
	return ctx
}
func (m *Monitor) listen(ctx context.Context) {
	defer m.cancel()
	// create a new ticker channel to listen to
	tick := time.NewTicker(time.Millisecond * 10)
	defer tick.Stop()
	// use an infinite loop to continue to listen
	// to new messages after the select statement
	// has been executed
	for {
		select {
		case <-ctx.Done(): // listen for context cancellation
			// shut down if the context is canceled
			fmt.Println("shutting down monitor")
			// if the monitor was told to shut down
			// then it should call its cancel function
			// so the client will know that the monitor
			// has properly shut down.
			m.cancel()
			// return from the function
			return
		case <-tick.C: // listen to the ticker channel
			// and print a message every time it ticks
			fmt.Println("monitor check")
		}
	}
}
