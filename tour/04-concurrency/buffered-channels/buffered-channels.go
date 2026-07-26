package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 - all goroutines are asleep - deadlock! it just clogs it
	// and panics
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
