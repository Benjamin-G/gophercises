package chaptertwelve

import "testing"

var global int64

func BenchmarkSum2(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum2(s)
	}
	global = local
}

func BenchmarkSum8(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum8(s)
	}
	global = local
}

const n = 1_000_000

func BenchmarkSumFoo(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]Foo, n)
		b.StartTimer()
		local = sumFoo(s)
	}
	global = local
}

func BenchmarkSumBar(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bar := Bar{
			a: make([]int64, n),
			b: make([]int64, n),
		}
		b.StartTimer()
		local = sumBar(bar)
	}
	global = local
}

var globalResult1 Result1

func BenchmarkCount1(b *testing.B) {
	var local Result1
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = count1(inputs)
	}
	globalResult1 = local
}

var globalResult2 Result2

func BenchmarkCount2(b *testing.B) {
	var local Result2
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = count2(inputs)
	}
	globalResult2 = local
}
