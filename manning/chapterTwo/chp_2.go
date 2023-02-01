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

type IntConfig struct {
	value int
	configName
	t    configType
	port int
}

func (c *IntConfig) GetConfigTypeDetails() string {
	return c.t.Get()
}

// builder pattern

type ConfigBuilder struct {
	value *int
	port  *int
	t     *configType
}

func (b *ConfigBuilder) Value(value int) *ConfigBuilder {
	b.value = &value
	return b
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) ConfigName(cN configType) *ConfigBuilder {
	b.t = &cN
	return b
}

func (b *ConfigBuilder) Build() (IntConfig, error) {
	cfg := IntConfig{}

	if b.port == nil {
		cfg.port = 8080
	} else {
		cfg.port = *b.port
	}

	if b.value == nil {
		cfg.value = 0
	} else {
		cfg.value = *b.value
	}

	if b.t == nil {
		cfg.t = configType{details: "Default Config"}
	} else {
		cfg.t = *b.t
	}

	return cfg, nil
}

// Restricting
type configName struct{ name string }

func (c *configName) GetName() string {
	return c.name
}

type configType struct{ details string }

func (c *configType) Get() string {
	return c.details
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

	// Embedded Structs
	// Promoted
	bar.configName = configName{name: "embedded"}
	fmt.Println(bar.configName.name)
	fmt.Println(bar.name)
	fmt.Println(bar.GetName())

	// not embedded
	// can create access to methods
	bar.t = configType{details: "Should not be embedded"}
	fmt.Println(bar.t.details)
	fmt.Println(bar.GetConfigTypeDetails())

	// builder pattern
	builder := ConfigBuilder{}

	cfg, _ := builder.Build()
	fmt.Println(cfg)
	fmt.Println(cfg.GetConfigTypeDetails())
	builder.Port(8000)
	builder.ConfigName(configType{details: "Should not be embedded"})
	builder.Value(1000)

	cfg, _ = builder.Build()
	fmt.Println(cfg)
	fmt.Println(cfg.GetConfigTypeDetails())

}

// Generics
// This interface is used by different functions such as sort.Ints or sort
// .Float64s.
type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }
