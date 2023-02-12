package chaptersix

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

func ChapterSixRunner() {
	pointer()
	nilRun()
	args()
}

func pointer() {
	c := customer{balance: 100.0}
	c.add(50.0)
	fmt.Printf("balance: %.2f\n", c.balance)
	c.add(50.0)
	fmt.Printf("balance: %.2f\n", c.balance)
	c.add(50.0)
	fmt.Printf("balance: %.2f\n", c.balance)
}

type customer struct {
	balance float64
}

func (c *customer) add(operation float64) {
	c.balance += operation
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

func nilRun() {
	customer := Customer{Name: "John", Age: 0}
	if err := customer.Validate1(); err != nil {
		log.Printf("customer is invalid: %v\n", err)
	}
	if err := customer.Validate2(); err != nil {
		log.Printf("customer is invalid: %v", err)
	}
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate1() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	return m
}

func (c Customer) Validate2() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	if m != nil {
		return m
	}
	return nil
}

func countEmptyLines(reader io.Reader) (int, error) {
	//file, err := os.Open(filename)
	//if err != nil {
	//	return 0, err
	//}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// ...
	}
	return 0, nil
}

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func args() {
	_ = f1()
	_ = f2()
	_ = f3()
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j) // prints 0 1
	}(i)
	i++
	j++
}

func f1() error {
	var status string
	defer notify(status)
	defer incrementCounter(status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func f2() error {
	var status string
	defer notifyPtr(&status)
	defer incrementCounterPtr(&status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func f3() error {
	var status string
	defer func() {
		notify(status)
		incrementCounter(status)
	}()

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func notify(status string) {
	fmt.Println("notify:", status)
}

func incrementCounter(status string) {
	fmt.Println("increment:", status)
}

func notifyPtr(status *string) {
	fmt.Println("notify:", *status)
}

func incrementCounterPtr(status *string) {
	fmt.Println("increment:", *status)
}

func foo() error {
	return nil
}

func bar() error {
	return nil
}
