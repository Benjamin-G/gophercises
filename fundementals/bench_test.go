package fundementals

import (
	"sync/atomic"
	"testing"
)

const rows = 1000

var res int64

var g int // Define a global variable

func BenchmarkAddTen(b *testing.B) {
	var v int
	for i := 0; i < b.N; i++ {
		v = addTen(1)
	}
	g = v
}

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}

func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}

var global uint64 // Define a global variable

func BenchmarkPopcnt2(b *testing.B) {
	var v uint64 // Define a local variable
	for i := 0; i < b.N; i++ {
		v = popcnt(uint64(i)) // Assign the result to the local variable
	}
	global = v // Assign the result to the global variable
}

// Coming back to the benchmark, the main issue is that we keep reusing the same matrix in both cases.
func BenchmarkCalculateSum512(b *testing.B) {
	var sum int64
	s := createMatrix512(rows) // Create a matrix of 512 columns
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum512(s) // Create a matrix of 512 columns
	}
	res = sum
}

func BenchmarkCalculateSum513(b *testing.B) {
	var sum int64
	s := createMatrix513(rows) // Create a matrix of 513 columns
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum513(s) // Calculate the sum
	}
	res = sum
}

func BenchmarkCalculateSum512_Accurate(b *testing.B) {
	var sum int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := createMatrix512(rows) // Create a new matrix during each loop iteration
		sum = calculateSum512(s)
	}
	res = sum
}

func BenchmarkCalculateSum513_Accurate(b *testing.B) {
	var sum int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := createMatrix513(rows) // Create a new matrix during each loop iteration
		sum = calculateSum513(s)
	}
	res = sum
}
