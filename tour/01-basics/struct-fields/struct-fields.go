package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("v = (%d,%d)\n", v.X, v.Y)
	fmt.Println("Scaling by 4")
	v.X = 4
	v.Y *= 4
	fmt.Printf("v = (%d,%d)\n", v.X, v.Y)
}
