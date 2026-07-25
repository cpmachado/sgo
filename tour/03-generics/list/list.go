package main

import (
	"fmt"
	"strings"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// stringer interface
func (l *List[T]) String() string {
	if l == nil {
		return "nil"
	}
	sb := strings.Builder{}

	sb.WriteString("(")
	ptr := l
	for ptr != nil {
		fmt.Fprintf(&sb, "%v", ptr.val)
		if ptr.next != nil {
			sb.WriteString(" ")
		}
		ptr = ptr.next
	}
	sb.WriteString(")")

	return sb.String()
}

func (l *List[T]) Len() int {
	ret := 0

	for l != nil {
		l = l.next
		ret++
	}

	return ret
}

func (l *List[T]) Map(f func(T) T) *List[T] {
	var ret, ptr *List[T]

	for l != nil {
		next := new(List[T])
		next.val = f(l.val)

		if ret == nil {
			ret = next
		} else {
			ptr.next = next
		}
		ptr = next
		l = l.next
	}

	return ret
}

func main() {
	n := 10
	a := newIntList(1, n)

	fmt.Println("a:", a)
	fmt.Printf("Len: %d\n", a.Len())

	a2 := a.Map(func(x int) int { return 2 * x })

	fmt.Println("2 * a:", a2)
	fmt.Printf("Len: %d\n", a2.Len())
}

func newIntList(a, b int) *List[int] {
	var ret, ptr *List[int]
	for a <= b {
		next := new(List[int])
		next.val = a

		if ret == nil {
			ret = next
		} else {
			ptr.next = next
		}
		ptr = next
		a++
	}
	return ret
}
