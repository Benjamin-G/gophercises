package exercises

import (
	"fmt"
	"sync"
)

// So, a mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.

// need to use a mutex to guard any updates to any variables.
// Everything you can achieve with a Mutex can be done with a channel in Go if the size of the channel is set to 1.

var (
	mutex   sync.Mutex
	balance int
)

// the use case for what is known as a binary semaphore - a semaphore/channel of size 1 - is so common in the real world that it made sense to implement this exclusively in the form of a mutex

func init() {
	balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func Mutex() {
	fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
