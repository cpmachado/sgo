package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// the cast to float64 is required, or the fmt function will find it's an
	// error and call Error() over it, inducing an infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return sqrtWithInitial(x, 1), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
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
