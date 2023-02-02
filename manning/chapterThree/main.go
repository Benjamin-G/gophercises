package chapterThree

import (
	"fmt"
	"math"
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
