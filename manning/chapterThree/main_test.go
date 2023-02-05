package chapterThree

import (
	"math"
	"math/rand"
	"testing"
)

func Test_IntOverflow(t *testing.T) {
	t.Parallel()
	t.Run("Inc32", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		var i int32 = math.MaxInt32
		Inc32(i)
	})

	t.Run("IncInt", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		i := math.MaxInt
		IncInt(i)
	})

	t.Run("IncUint", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		var i uint = math.MaxUint
		IncUint(i)
	})

	t.Run("AddInt", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		i := math.MaxInt
		AddInt(i, i)
	})

	t.Run("MultiplyInt", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		i := math.MaxInt
		MultiplyInt(i, i)
	})
}

func createRandomIntArray(t testing.TB, n int) []int {
	t.Helper()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(50000000))
	}
	return arr
} // Define a global variable
// BenchmarkConvertToFloat64-8         6036            242939 ns/op
func BenchmarkConvertToFloat64(b *testing.B) {
	testArr := createRandomIntArray(b, 1_000_000)
	for i := 0; i < b.N; i++ {
		convertToFloat64(testArr)
	}
}

// BenchmarkConvertToFloat64Better-8          21009             50575 ns/op
func BenchmarkConvertToFloat64GivenCapacity(b *testing.B) {
	testArr := createRandomIntArray(b, 1_000_000)
	for i := 0; i < b.N; i++ {
		convertToFloat64GivenCapacity(testArr)
	}
}

func BenchmarkConvertToFloat64GivenLength(b *testing.B) {
	testArr := createRandomIntArray(b, 1_000_000)
	for i := 0; i < b.N; i++ {
		convertToFloat64GivenLength(testArr)
	}
}

const n = 1_000_000

var globalBar []Bar

func BenchmarkConvert_EmptySlice(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertEmptySlice(foos)
	}
	globalBar = local
}

func BenchmarkConvert_GivenCapacity(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenCapacity(foos)
	}
	globalBar = local
}

func BenchmarkConvert_GivenLength(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenLength(foos)
	}
	globalBar = local
}

var global map[int]struct{}

func BenchmarkMapWithoutSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{})
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}

// 60% faster
func BenchmarkMapWithSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{}, n)
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}
