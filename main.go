package main

import (
	"fmt"

	"github.com/kweheliye/golang-playground/types"
)

func main() {

	//types examples
	fmt.Printf("Has a: %v\n", types.Has([]string{"a", "b"}, "a"))
	fmt.Printf("Has c: %v\n", types.Has([]string{"a", "b"}, "c"))

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		types.SumIntsOrFloats[string, int64](ints),
		types.SumIntsOrFloats[string, float64](floats))
}
