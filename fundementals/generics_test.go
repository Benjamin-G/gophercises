package fundementals

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Keys(t *testing.T) {
	t.Parallel()
	// create a map with some values
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// create an interstitial map to pass to the function
	im := map[any]any{}
	// copy the map into the interstitial map
	for k, v := range m {
		im[k] = v
	}

	// get the keys
	keys := Keys(im)
	// create a slice to hold the keys as
	// integers for comparison
	act := make([]int, 0, len(keys))
	// copy the keys into the integer slice
	for _, k := range keys {
		// assert that the key is an int
		i, ok := k.(int)
		if !ok {
			t.Fatalf("expected type int, got %T", k)
		}
		act = append(act, i)
	}
	// sort the returned keys for comparison
	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})
	// set the expected values
	exp := []int{1, 2, 3}

	// assert the length of the actual and expected values
	al := len(act)
	el := len(exp)

	if al != el {
		t.Fatalf("expected %d, but got %d", el, al)
	}
	// loop through the expected values and
	// assert they are in the actual values
	for i, v := range exp {
		if v != act[i] {
			t.Fatalf("expected %d, but got %d", v, act[i])
		}
	}
}

func Test_KeysGeneric(t *testing.T) {
	t.Parallel()
	// create a map with some values
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	// create a function variable pointing
	// to the Keys function
	fn := KeysGeneric[int, string]
	// get the keys
	act := fn(m)
	// sort the returned keys for comparison
	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})
	// set the expected values
	exp := []int{1, 2, 3}

	// // set the expected values
	// exp := []string{"one", "two", "three"}

	// assert the length of the actual and expected values
	if len(exp) != len(act) {
		t.Fatalf("expected len(%d), but got len(%d)", len(exp), len(act))
	}
	// assert the types of the actual and expected values
	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)
	if at != et {
		t.Fatalf("expected type %s, but got type %s", et, at)
	}
	// loop through the expected values and
	// assert they are in the actual values
	for i, v := range exp {
		if v != act[i] {
			t.Fatalf("expected %d, but got %d", v, act[i])
		}
	}
}

func Test_KeysGeneric_MyInt(t *testing.T) {
	t.Parallel()
	// create a map with some values
	m := map[MyInt]string{
		1: "one",
		2: "two",
		3: "three",
	}
	// create a function variable pointing
	// to the Keys function
	fn := KeysGeneric[MyInt, string]
	// get the keys
	act := fn(m)
	// sort the returned keys for comparison
	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})
	// set the expected values
	exp := []MyInt{1, 2, 3}

	// // set the expected values
	// exp := []string{"one", "two", "three"}

	// assert the length of the actual and expected values
	if len(exp) != len(act) {
		t.Fatalf("expected len(%d), but got len(%d)", len(exp), len(act))
	}
	// assert the types of the actual and expected values
	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)
	if at != et {
		t.Fatalf("expected type %s, but got type %s", et, at)
	}
	// loop through the expected values and
	// assert they are in the actual values
	for i, v := range exp {
		if v != act[i] {
			t.Fatalf("expected %d, but got %d", v, act[i])
		}
	}
}

func Test_KeysGeneric_Float(t *testing.T) {
	t.Parallel()
	// create a map with some values
	m := map[float64]string{
		1.1: "one",
		2.2: "two",
		3.3: "three",
	}
	// create a function variable pointing
	// to the Keys function
	fn := KeysGeneric[float64, string]
	// get the keys
	act := fn(m)
	// sort the returned keys for comparison
	sort.Slice(act, func(i, j int) bool {
		return act[i] < act[j]
	})
	// set the expected values
	exp := []float64{1.1, 2.2, 3.3}

	// // set the expected values
	// exp := []string{"one", "two", "three"}

	// assert the length of the actual and expected values
	if len(exp) != len(act) {
		t.Fatalf("expected len(%d), but got len(%d)", len(exp), len(act))
	}
	// assert the types of the actual and expected values
	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)
	if at != et {
		t.Fatalf("expected type %s, but got type %s", et, at)
	}
	// loop through the expected values and
	// assert they are in the actual values
	for i, v := range exp {
		if v != act[i] {
			t.Fatalf("expected %v, but got %v", v, act[i])
		}
	}
}

func Test_Slicer(t *testing.T) {
	t.Parallel()
	// create input string
	input := "Hello World"
	// capture output []string
	act := Slicer(input)
	exp := []string{input}
	if len(act) != len(exp) {
		t.Fatalf("expected %v, got %v", exp, act)
	}
	for i, v := range exp {
		if act[i] != v {
			t.Fatalf("expected %v, got %v", exp, act)
		}
	}
}
