package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	// set i through pointer
	*p = 24
	fmt.Println(i)

	p = &j
	// set j and compute through pointer
	*p = *p / 37
	fmt.Println(j)

	// cpmachado: there's no pointer arithmetic unlike C, but I heard of
	// unsafe pointers, need to investigate
}
