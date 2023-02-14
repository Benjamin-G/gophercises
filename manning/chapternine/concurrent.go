package chapternine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func ConcurrentRunner() {
	loop()

	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{})

	go listing2(messageCh, disconnectCh)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}

	time.Sleep(10 * time.Millisecond)
}

func loop() {
	func() {
		s := []int{1, 2, 3}

		for _, i := range s {
			go func() {
				fmt.Print(i)
			}()
		}
	}()

	func() {
		s := []int{1, 2, 3}

		for _, i := range s {
			val := i
			go func() {
				fmt.Print(val)
			}()
		}
	}()

	func() {
		s := []int{1, 2, 3}

		for _, i := range s {
			go func(val int) {
				fmt.Print(val)
			}(i)
		}
	}()
}

func listing1(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func listing2(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}

func CustomerRunner() {
	customer := Customer{}

	err := customer.UpdateAge1(-1)
	if err != nil {
		fmt.Println(err)
	}

	err = customer.UpdateAge2(-1)
	if err != nil {
		fmt.Println(err)
	}

	err = customer.UpdateAge3(-1)
	if err != nil {
		fmt.Println(err)
	}
}

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

// How should we deal with this situation? First, it illustrates how unit
// testing is important.
func (c *Customer) UpdateAge1(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		//err := fmt.Errorf("age should be positive given: %v", c) will create a deadlock
		//Perhaps we want to call another function that doesn’t try to acquire the
		//mutex, or we only want to change the way we format the error so that it doesn’t call
		//the String method. For example, the following code doesn’t lead to a deadlock
		//because we only log the customer ID in accessing the id field directly:
		err := fmt.Errorf("age should be positive given: %d", c.age)
		return err
	}

	c.age = age
	return nil
}

func (c *Customer) UpdateAge2(age int) error {
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.age = age
	return nil
}

func (c *Customer) UpdateAge3(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer id %s", c.id)
	}

	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func MutexRunner() {
	c := Cache{
		balances: make(map[string]float64),
	}
	go c.AddBalance("1", 10.0)
	go c.AddBalance("2", 30.0)
	go c.AddBalance("3", 30.0)
	go c.AddBalance("4", 20.0)
	go c.AddBalance("5", 20.0)
	time.Sleep(10 * time.Millisecond)
	fmt.Println(c.AverageBalance1())
	fmt.Println(c.AverageBalance2())
	fmt.Println(c.AverageBalance3())

	//sync.WaitGroup{}
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)

	wg.Add(1)
	syncCond()
	wg.Wait()
	time.Sleep(10 * time.Millisecond)
}

func syncCond() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	max := 15

	go f(10)
	go f(max)

	// Updater goroutine
	for {
		time.Sleep(time.Second * (1 / 2))
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
		fmt.Println("broadcast", donation.balance)

		if donation.balance >= max {
			donation.cond.Broadcast()

			break
		}
	}
}

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

//This solution uses a sync.RWMutex to allow multiple readers as long as
//there are no writers.

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance1() float64 {
	c.mu.RLock()
	balances := c.balances

	//Internally, a map is a runtime.hmap struct containing mostly metadata (for example,
	//	a counter) and a pointer referencing data buckets. So, balances := c.balances
	//	doesn’t copy the actual data.
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}

	return sum / float64(len(balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()

	//If the iteration operation isn’t heavy (that’s the case here, as we perform an increment
	//operation), we should protect the whole function
	defer c.mu.RUnlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}

	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance3() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))

	for k, v := range c.balances {
		m[k] = v
	}

	//if the iteration operation isn’t lightweight, is to work on an actual
	//copy of the data and protect only the copy
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}

	return sum / float64(len(m))
}
