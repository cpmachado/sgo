package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println("The Beatles")
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("The Beatles first 2 duos")
	fmt.Println(a, b)

	fmt.Println("The Beatles erasing John")
	a[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println("The Beatles erasing Paul")
	b[0] = "YYY"
	fmt.Println(a, b)
	fmt.Println("The Beatles without John, and Paul")
	fmt.Println(names)
}
