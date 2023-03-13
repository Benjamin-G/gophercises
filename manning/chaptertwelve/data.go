package chaptertwelve

const n_ = 1_000_000

func add(s [2]int64) [2]int64 {
	for i := 0; i < n_; i++ {
		s[0]++
		if s[0]%2 == 0 {
			s[1]++
		}
	}

	return s
}

func add2(s [2]int64) [2]int64 {
	for i := 0; i < n_; i++ {
		v := s[0]
		s[0] = v + 1

		if v%2 != 0 {
			s[1]++
		}
	}

	return s
}

type Foo1 struct {
	b1 byte
	i  int64
	b2 byte
}

func sum1(foos []Foo1) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

type Foo2 struct {
	i  int64
	b1 byte
	b2 byte
}

func sum_2(foos []Foo2) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

// First, we use the println built-in function instead
// of fmt.Println, which would force allocating the c variable on the heap. Second, we
// disable inlining on the sumValue function; otherwise, the function call would not
// occur
func listing1() {
	a := 3
	b := 2

	c := sumValue(a, b)
	println(c)
}

//go:noinline
func sumValue(x, y int) int {
	z := x + y
	return z
}

func listing2() {
	a := 3
	b := 2

	c := sumPtr(a, b)
	println(*c)
}

//go:noinline
func sumPtr(x, y int) *int {
	z := x + y
	return &z
}

func listing3() {
	a := 3
	b := 2
	c := sum(&a, &b)
	println(c)
}

//go:noinline
func sum(x, y *int) int {
	return *x + *y
}
