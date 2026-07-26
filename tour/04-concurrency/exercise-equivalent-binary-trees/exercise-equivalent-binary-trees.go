package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for a := range ch1 {
		if a != <-ch2 {
			drainChannel(ch1)
			drainChannel(ch2)
			return false
		}
	}
	if len(ch2) > 0 {
		drainChannel(ch2)
		return false
	}

	return true
}

func drainChannel(ch chan int) {
	for range ch {
	}
}

func main() {
	ch := make(chan int)
	i := 0

	go Walk(tree.New(1), ch)

	fmt.Print("Walk(tree.New(1), ch): {")
	for c := range ch {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(c)
		i++
	}
	fmt.Println("}")
	fmt.Println("Same(tree.New(1), tree.New(1)) = ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)) = ", Same(tree.New(1), tree.New(2)))
}
