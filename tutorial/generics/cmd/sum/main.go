package main

import (
	"fmt"

	"sum"
)

func main() {
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

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sum.SumInts(ints),
		sum.SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n",
		sum.SumIntsOrFloats[string, int64](ints),
		sum.SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		sum.SumIntsOrFloats(ints),
		sum.SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		sum.SumNumbers(ints),
		sum.SumNumbers(floats))
}
