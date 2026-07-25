package main

import "fmt"

func main() {
	// nowadays, I normally set it to [any], it makes no difference
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
