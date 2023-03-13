package chaptertwelve

import "testing"

var global_ [2]int64

func BenchmarkAdd(b *testing.B) {
	a := [2]int64{}
	var local [2]int64
	for i := 0; i < b.N; i++ {
		local = add(a)
	}
	global_ = local
}

func BenchmarkAdd2(b *testing.B) {
	a := [2]int64{}
	var local [2]int64
	for i := 0; i < b.N; i++ {
		local = add2(a)
	}
	global_ = local
}

func BenchmarkSum1(b *testing.B) {
	var local int64
	s := make([]Foo1, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sum1(s)
	}
	global = local
}

func BenchmarkSum_2(b *testing.B) {
	var local int64
	s := make([]Foo2, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sum_2(s)
	}
	global = local
}
