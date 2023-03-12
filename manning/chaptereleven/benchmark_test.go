package chaptereleven

import "testing"

func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}

var global uint64

// It guarantees that we
// avoid the inlining optimizations, which can artificially lower the execution time or
// even remove the call to the function under test.
func BenchmarkPopcnt2(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = popcnt(uint64(i))
	}
	global = v
}

const rows = 1000

var res int64

func BenchmarkCalculateSum512_1(b *testing.B) {
	var sum int64
	s := createMatrix512(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum512(s)
	}
	res = sum
}

func BenchmarkCalculateSum513_1(b *testing.B) {
	var sum int64
	s := createMatrix513(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum513(s)
	}
	res = sum
}

func BenchmarkCalculateSum512_2(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix512(rows)
		b.StartTimer()
		sum = calculateSum512(s)
	}
	res = sum
}

func BenchmarkCalculateSum513_2(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix513(rows)
		b.StartTimer()
		sum = calculateSum513(s)
	}
	res = sum
}

func createMatrix512(r int) [][512]int64 {
	return make([][512]int64, r)
}

func createMatrix513(r int) [][513]int64 {
	return make([][513]int64, r)
}
