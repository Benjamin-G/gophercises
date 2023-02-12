// Package stringset provides basic constants and mathematical functions.
//
// This package does not guarantee bit-identical results
// across architectures.
package stringset

import "fmt"

// Set returns the fastest path between two points.
// Deprecated: This function uses a deprecated way to compute
// the fastest path. Use Set instead.
type Set map[string]struct{}

func New(...string) Set { return nil }

func (s Set) Sort() []string { return nil }

func SetRunner() {
	i := 0
	if true {
		var i = 1
		fmt.Println(i)
	}
	fmt.Println(i)
}
