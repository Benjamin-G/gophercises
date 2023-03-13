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
