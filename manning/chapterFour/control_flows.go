package chapterFour

import (
	"context"
	"fmt"
)

func ControlRunner() {
	arrays()
	channels()
	pointer()
	fmt.Printf("listing1\n")
	listing1()
	fmt.Printf("\nlisting2\n")
	listing2()
}

type account struct {
	balance float32
}

func arrays() {
	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	// when []*account this will mutate the slice
	//for _, a := range accounts {
	//	a.balance += 1000
	//}
	//for _, a := range accounts {
	//	fmt.Println(a)
	//}
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	s := []int{0, 1, 2}
	for range s {
		s = append(s, 10)
	}
	fmt.Println(s)
}
func channels() {
	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		close(ch2)
	}()

	ch := ch1
	for v := range ch {
		fmt.Println(v)
		ch = ch2
	}
}

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func pointer() {
	s := Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
	s.storeCustomers2([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
	s.storeCustomers3([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func (s *Store) storeCustomers2(customers []Customer) {
	for _, customer := range customers {
		current := customer
		s.m[current.ID] = &current
	}
}

func (s *Store) storeCustomers3(customers []Customer) {
	for i := range customers {
		customer := &customers[i]
		s.m[customer.ID] = customer
	}
}

func print(m map[string]*Customer) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%#v\n", k, v)
	}
}

func listingMap1() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}

	fmt.Println(m)
}

func listingMap2() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)
	//m2 := map[int]bool{}

	for k, v := range m {
		//m2[k] = v
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	res := make(map[int]bool, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

func listing1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 1:
			fmt.Println("One")
		case 2:
			//return
			break
		case 3:
			fmt.Println("Three")
		}

	}
}

func listing2() {
numLoop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break numLoop
		}
	}
}

func listing3(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break
		}
	}
}

func listing4(ctx context.Context, ch <-chan int) {
loop:
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break loop
		}
	}
}
