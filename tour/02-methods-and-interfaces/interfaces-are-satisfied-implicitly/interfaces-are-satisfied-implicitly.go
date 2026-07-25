package main

import "fmt"

type I interface {
	M()
}

// No need to explicitly say it complies with the interface, as it is inferred
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
