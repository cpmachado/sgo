package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Edge struct {
	A *Vertex
	B *Vertex
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Edge{A: &Vertex{1, 2}, B: &Vertex{2, 4}})
}
