package fundementals

import "golang.org/x/exp/constraints"

func Keys(m map[any]any) []any {
	// make a slice of the keys
	keys := make([]any, 0, len(m))
	// iterate over the map
	for k := range m {
		// add the key to the slice
		keys = append(keys, k)
	}
	// return the keys
	return keys
}

type MapKey interface {
	~int | ~float64
}

// MapKey is a set of a constraints
// on types that can be used as map keys.
// type MapKeyInt interface {
// 	~int
// }

type MyInt int

// K int | float64 or K MapKey or constraints.Ordered

func KeysGeneric[K constraints.Ordered, P any](m map[K]P) []K {
	// snippet: def
	// make a slice of the keys
	keys := make([]K, 0, len(m))
	// iterate over the map
	for k := range m {
		// add the key to the slice
		keys = append(keys, k)
	}
	// return the keys
	return keys
}

func Slicer[T any](input T) []T {
	return []T{input}
}

func SlicerStrings(input string) []string {
	return []string{input}
}
