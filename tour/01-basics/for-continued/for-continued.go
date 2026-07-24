package main

import "fmt"

func main() {
	sum := 1

	// cpmachado: doesn't require semicolons when it's just a condition
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
