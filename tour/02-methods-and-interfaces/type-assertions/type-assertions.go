package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// avoid panic, by using previous form, don't assume what you don't know
	// f = i.(float64) // panic
	// fmt.Println(f)
}
