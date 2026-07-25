package main

import "fmt"

type A struct {
	X, Y int
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// cpmachado: you can do it with structs, because structs of comparable
	// types are comparable
	sa := []A{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}}
	fmt.Println(Index(sa, A{X: 1, Y: 2}))
}
