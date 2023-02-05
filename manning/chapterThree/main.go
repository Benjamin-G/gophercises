package chapterThree

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func MathsRunner() {
	//100 + base 8 10
	sum := 100 + 010
	fmt.Println(sum)

	//1.0001 * 1.0001 = 1.0002
	var n float32 = 1.0001
	fmt.Println(n * n)

	//1.0001 * 1.0001 = 1.00020001 float64
	m := 1.0001
	fmt.Println(m * m)

	fmt.Println(math.SmallestNonzeroFloat64)

	// special case floating points
	var a float64
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Println(positiveInf, negativeInf, nan)
	z := 1 / a
	fmt.Println(z)

	f1 := func(n int) float64 {
		result := 10_000.
		for i := 0; i < n; i++ {
			result += 1.0001
		}
		return result
	}

	f2 := func(n int) float64 {
		result := 0.
		for i := 0; i < n; i++ {
			result += 1.0001
		}
		return result + 10_000.
	}

	// 10010.001
	fmt.Println(f1(10))
	fmt.Println(f2(10))

	{
		a := 100000.001
		b := 1.0001
		c := 1.0002
		//200030.00200030004
		//200030.0020003
		fmt.Println(a * (b + c))
		fmt.Println(a*b + a*c)
		//The exact result is 200,030.002.
	}

	// Slices
	fmt.Println("Slices:")
	{
		s := make([]int, 3, 6)
		println(cap(s))
		println(len(s))
		s = append(s, 0, 1, 0)
		s = append(s, 0, 1, 0)
		s = append(s, 2, 2, 2)
		fmt.Println(s)
		println(cap(s))
		println(len(s))
	}
	{
		s := []int{}
		println(cap(s))
		println(len(s))
		s = append(s, 0, 1, 0)
		s = append(s, 0, 1, 0)
		s = append(s, 2, 2, 2)
		fmt.Println(s)
		println(cap(s))
		println(len(s))
	}
	{
		s1 := make([]int, 3, 6)
		s1[1] = 1
		//s2 := s1[1:3]
		s2 := append(s1[1:3], 2)
		s2 = append(s2, 3)
		s2 = append(s2, 4)
		s2 = append(s2, 5)
		fmt.Println(s2)
	}
}

func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow")
	}
	return counter + 1
}

func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("int overflow")
	}
	return counter + 1
}

func IncUint(counter uint) uint {
	if counter == math.MaxUint {
		panic("uint overflow")
	}
	return counter + 1
}

func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("int overflow")
	}

	return a + b
}

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt || b == math.MinInt {
		panic("integer overflow")
	}
	if result/b != a {
		panic("integer overflow")
	}
	return result
}

// performance
func convertToFloat64(arr []int) []float64 {
	//res := make([]float64, 0)
	//var res []float64
	res := []float64(nil)
	for _, n := range arr {
		res = append(res, float64(n))
	}
	return res
}

func convertToFloat64GivenCapacity(arr []int) []float64 {
	res := make([]float64, 0, len(arr))
	for _, n := range arr {
		res = append(res, float64(n))
	}
	return res
}

// 4% faster might be less readable
func convertToFloat64GivenLength(arr []int) []float64 {
	res := make([]float64, len(arr))
	for i, n := range arr {
		res[i] = float64(n)
	}
	return res
}

func convertEmptySlice(foos []Foo) []Bar {
	bars := make([]Bar, 0)

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenCapacity(foos []Foo) []Bar {
	n := len(foos)
	bars := make([]Bar, 0, n)

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenLength(foos []Foo) []Bar {
	n := len(foos)
	bars := make([]Bar, n)

	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}

type Foo struct{}

type Bar struct{}

func fooToBar(foo Foo) Bar {
	return Bar{}
}

func SliceRunner() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b))

	// Copy Slice
	src := []int{0, 1, 2}
	dst := append([]int(nil), src...)
	fmt.Println(src)
	fmt.Println(dst)

	{
		s1 := []int{1, 2, 3}
		s2 := s1[:2]
		// Mutates s1
		s3 := append(s2, 10)
		fmt.Printf("%v, %v, %v\n", cap(s1), cap(s2), cap(s3))
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		printAlloc()
	}
	{
		s1 := []int{1, 2, 3}
		s2 := append([]int(nil), s1[1:3]...)
		// Mutates s1
		s3 := append(s2, 10)
		fmt.Printf("%v, %v, %v\n", cap(s1), cap(s2), cap(s3))
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		printAlloc()
	}

	fmt.Println("Memory and pointers")
	foos := make([]Foo2, 1_000)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo2{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()
	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

type Foo2 struct {
	v []byte
}

func keepFirstTwoElementsOnly(foos []Foo2) []Foo2 {
	// Bad
	// return foos[:2]

	// Good
	//res := make([]Foo2, 2)
	//copy(res, foos)
	//return res

	// Slightly better
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB ", m.Alloc/1024)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

type customer struct {
	ID         string
	Operations []float32
}

func MemoryLeakRunner() {
	// Init
	n := 1_000_000
	{ //m := make(map[int][128]byte)
		m := make(map[int]*[128]byte)
		//printAlloc()
		//m = make(map[int][128]byte, n)
		printAlloc()

		// Add elements
		for i := 0; i < n; i++ {
			m[i] = randBytes()
		}
		printAlloc()

		// Remove elements
		for i := 0; i < n; i++ {
			delete(m, i)
		}

		// End
		runtime.GC()
		printAlloc()
		runtime.KeepAlive(m)
	}
	{
		// This will autmatically use pointers
		m := make(map[int][256]byte)
		//printAlloc()
		printAlloc()

		// Add elements
		for i := 0; i < n; i++ {
			m[i] = randBytes256()
		}
		printAlloc()

		// Remove elements
		for i := 0; i < n; i++ {
			delete(m, i)
		}

		// End
		runtime.GC()
		printAlloc()
		runtime.KeepAlive(m)
	}
}

func randBytes() *[128]byte {
	return &[128]byte{}
}

func randBytes256() [256]byte {
	return [256]byte{}
}

type customer1 struct {
	id string
}

type customer2 struct {
	id         string
	operations []float64
}

func (a customer2) equal(b customer2) bool {
	if a.id != b.id {
		return false
	}
	if len(a.operations) != len(b.operations) {
		return false
	}
	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}

func CompareRunner() {
	cust11 := customer1{id: "x"}
	cust12 := customer1{id: "x"}
	fmt.Println(cust11 == cust12)
	//Doing a few benchmarks locally with structs of
	//different sizes, on average, reflect.DeepEqual is about 100 times slower than ==.
	fmt.Println(reflect.DeepEqual(cust11, cust12))

	cust21 := customer2{id: "x", operations: []float64{1.}}
	cust22 := customer2{id: "x", operations: []float64{1.}}
	// Doesn't compile
	// fmt.Println(cust21 == cust22)
	_ = cust21
	_ = cust22

	var a any = 3
	var b any = 3
	fmt.Println(a == b)

	var cust31 = customer2{id: "x", operations: []float64{1.}}
	var cust32 = customer2{id: "x", operations: []float64{1.}}
	fmt.Println(cust31.equal(cust32))

	cust41 := customer2{id: "x", operations: []float64{1.}}
	cust42 := customer2{id: "x", operations: []float64{1.}}
	fmt.Println(reflect.DeepEqual(cust41, cust42))
}
