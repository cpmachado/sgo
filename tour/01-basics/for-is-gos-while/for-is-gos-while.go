package main

import "fmt"

func main() {
	sum := 1
	// cpmachado: basically the same as for continued, whence I removed the
	// semicolons
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
