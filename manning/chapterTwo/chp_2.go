package chapterTwo

import (
	"fmt"
	"io"
	"log"
	"sort"
)

type ErrJoin struct {
	err error
}

func (e ErrJoin) Error() string {
	return fmt.Sprintf("[ErrJoin] %s", e.err)
}

type ErrConcat struct {
	err error
}

func (e ErrConcat) Error() string {
	return fmt.Sprintf("[ErrConcat] %s", e.err)
}

func NestedRunner() {
	s, err := join2("chapter", "hello world", 5)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}

// Code smell
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", ErrJoin{err: fmt.Errorf("s1 is empty")}
	} else {
		if s2 == "" {
			return "", ErrJoin{err: fmt.Errorf("s2 is empty")}
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", ErrJoin{err: fmt.Errorf("s1 is empty")}
	}
	if s2 == "" {
		return "", ErrJoin{err: fmt.Errorf("s2 is empty")}
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		//concat = concat[:max]
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", ErrConcat{err: fmt.Errorf("concatenate")}
	//return s1 + " " + s2, nil
}

func copySourceToDest(source io.Reader, dest io.Writer) error {
	b, err := io.ReadAll(source)
	if err != nil {
		return err
	}
	_, err = dest.Write(b)
	return err
}

// SORT Example
type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

type Organ struct {
	Name   string
	Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

// ByWeight implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func isSortedAsc(data sort.Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func isSortedDesc(data sort.Interface) bool {
	//n := data.Len()
	for i := 0; i > data.Len()-1; i++ {
		if data.Less(i, i+1) {
			return false
		}
	}
	return true
}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}

func SortRunner() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight:")
	printOrgans(s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)

	sort.Slice(s, func(i, j int) bool { return s[i].Weight > s[j].Weight })
	fmt.Println("Organs descending by weight:")
	printOrgans(s)
	//sort.Slice(people, func(i, j int) bool {
	//	return people[i].Age > people[j].Age
	//})
	//fmt.Println("Is sorted desc: ", isSortedDesc(ByAge(people)))
	//fmt.Println(people)
}

// Restricting
type IntConfig struct {
	value int
}

func (c *IntConfig) Get() int {
	return c.value
}

func (c *IntConfig) Set(value int) {
	c.value = value
}

func (c *IntConfig) Double() {
	c.value = c.value * 2
}

type doubler interface {
	Double()
}

func dub(c doubler) {
	c.Double()
}

type intConfigGetter interface {
	Get() int
}

type Foo struct {
	threshold intConfigGetter
}

func NewFoo(threshold intConfigGetter) Foo {
	return Foo{threshold: threshold}
}

func (f Foo) Bar() {
	threshold := f.threshold.Get()
	fmt.Println(threshold)
}

func RestrictionRunner() {
	foo := NewFoo(&IntConfig{value: 42})
	if _, ok := foo.threshold.(intConfigGetter); ok {
		foo.Bar()
	}
	// Has all functionality of IntConfig
	bar := IntConfig{value: 24}
	fmt.Println(bar.Get())
	bar.Set(42)
	fmt.Println(bar.Get())
	bar.Double()
	fmt.Println(bar.Get())
	dub(&bar)
	fmt.Println(bar.Get())
}
