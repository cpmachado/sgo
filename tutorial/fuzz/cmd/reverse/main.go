package main

import (
	"fmt"
	"reverse"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := reverse.Reverse(input)
	doubleRev, doubleRevErr := reverse.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
