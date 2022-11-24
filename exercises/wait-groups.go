package exercises

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// https://tutorialedge.net/golang/go-waitgroup-tutorial/

// myFunc passes lock by value: sync.WaitGroup contains sync.noCopycopylocks
func myFunc(waitgroup *sync.WaitGroup, name string, num int) {
	if num != 0 {
		num = rand.Intn(10)
	}
	waitgroup.Add(1)
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Inside my goroutine", name, num)
	waitgroup.Done()
}

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	wg.Add(1)
	time.Sleep(2000 * time.Millisecond)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	// fmt.Println(resp.Status)
	return resp.Status, nil
}

func WaitGroups() {
	fmt.Println("WaitGroups Starting...")
	var waitgroup sync.WaitGroup
	sumNum := 10
	go myFunc(&waitgroup, "1", sumNum)
	// panic: sync: negative WaitGroup counter if ran without add
	go myFunc(&waitgroup, "2", sumNum)

	// Cannot add to WaitGroup inside anonymous func
	// go func(waitgroup *sync.WaitGroup) {
	// 	waitgroup.Add(1)
	// 	fmt.Println("Anonymous Functions, Inside my goroutine")
	// 	waitgroup.Done()
	// }(&waitgroup)
	numbah := 9
	waitgroup.Add(1)
	go func() {
		fmt.Println("Anonymous Functions, Inside my goroutine", numbah, sumNum)
		waitgroup.Done()
	}()

	waitgroup.Wait()

	fmt.Println("Executes before the For loop")

	// It appears that the for loop blocks until it reaches completion
	for _, url := range urls {
		s, err := fetch(url, &waitgroup)
		if err == nil {
			fmt.Println(url, " :: ", s)
		} else {
			// if there was an error we can count this as "completed"
			waitgroup.Done()
		}
	}

	fmt.Println("Before wait!")
	// Adding too many to your waitgroup will block the goroutine
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}
