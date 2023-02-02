package chapterThree

import (
	"math"
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
