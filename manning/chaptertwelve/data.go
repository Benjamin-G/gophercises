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
