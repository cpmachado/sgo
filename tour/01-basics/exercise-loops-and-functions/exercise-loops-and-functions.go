package main

import (
	"fmt"
	"math"
	"slices"
)

func Sqrt(x float64) float64 {
	return sqrtWithInitial(x, 1)
}

func main() {
	x := float64(2)
	// slices, I know this isn't covered yet
	xis := []float64{1, 3 * x / 4, 3 * x / 2, x}

	// n = 10 was original
	n := 30
	// epsilon, considering value <= 10^{-15} is irrelevant
	eps := math.Pow(10, -15)

	// sort initial values
	slices.Sort(xis)

	for _, xi := range xis {
		res, k := sqrtBoundWithInitialAndEps(x, xi, n, eps)
		diff := math.Sqrt(x) - res
		fmt.Printf("(xi, res, k, diff) = (%g, %g, %v, %g)\n", xi, res, k, diff)
	}
}

func sqrtWithInitial(x float64, initial float64) float64 {
	z, _ := sqrtWithInitialAndEps(x, initial, 0)
	return z
}

func sqrtWithInitialAndEps(x float64, initial float64, eps float64) (float64, int) {
	return sqrtBoundWithInitialAndEps(x, initial, 10, eps)
}

func sqrtBoundWithInitialAndEps(x float64, initial float64, n int, eps float64) (float64, int) {
	z := initial

	for i := range n {
		d := (z*z - x) / (2 * z)
		z -= d
		if math.Abs(d) < eps {
			return z, i
		}
	}

	return z, n
}
