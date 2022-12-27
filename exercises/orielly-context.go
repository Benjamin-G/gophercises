package exercises

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Context is for
// To provide an API for canceling branches of your call-graph.
// To provide a data-bag for transporting request-scoped data through your call-graph.

// As we learned in “Preventing Goroutine Leaks”, cancellation in a function has three aspects:
// A goroutine’s parent may want to cancel it.
// A goroutine may want to cancel its children.
// Any blocking operations within a goroutine need to be preemptable so that it may be canceled.

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

// Although the difference in this iteration of the program is small, it allows the locale function to fail fast.
var reqID = 0

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		// we take the deadline and check if this function will succeed, line  in select
		if deadline.Sub(time.Now().Add(5*time.Second)) <= 0 {
			return "", errors.New("fast fail")
		}
	}

	handleResponse(ctx)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(5 * time.Second):
	}
	return "EN/US", nil
}

func processRequest(ctx context.Context, userID, authToken string) context.Context {
	// The data should help decorate operations, not drive them.
	// The data should be immutable.
	// Use context values only for request-scoped data that transits processes and
	// API boundaries, not for passing optional parameters to functions.
	ctx = context.WithValue(ctx, ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	ctx = context.WithValue(ctx, ctxCallCount, reqID)
	handleResponse(ctx)
	return ctx
}

func handleResponse(ctx context.Context) {
	reqID += 1
	fmt.Printf(
		"handling response for %v (%v) called(%v,%v)\n",
		userID(ctx),
		authToken(ctx),
		reqID,
		callCount(ctx),
	)
}

type foo int
type bar int

// First, they recommend you define a custom key-type in your package. As long as other packages do the same, this prevents collisions within the Context.

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
	ctxCallCount
)

func userID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func authToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func callCount(c context.Context) int {
	res, _ := strconv.Atoi(c.Value(ctxAuthToken).(string))
	return res
}

// If there are a few frameworks and tens of functions between where the data is accepted and where it is used, do you want to lean toward verbose, self-documenting function signatures, and add the data as a parameter? Or would you rather place it in a Context and thereby create an invisible dependency? There are merits to each approach, and in the end it’s a decision you and your team will have to make.

func OriellyContext() {
	fmt.Println("OriellyContext Starting...")
	var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2

	fmt.Printf("%v\n", m)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx = processRequest(ctx, "jane", "abc123")
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}
