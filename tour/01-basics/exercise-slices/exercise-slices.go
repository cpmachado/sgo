package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	f := Xor
	return picf(dx, dy, f)
}

func main() {
	pic.Show(Pic)
}

func picf(dx, dy int, f func(uint8, uint8) uint8) [][]uint8 {
	res := make([][]uint8, dy)

	for y := range dy {
		res[y] = make([]uint8, dx)
		for x := range dx {
			res[y][x] = f(uint8(x), uint8(y))
		}
	}
	return res
}

func Average(x, y uint8) uint8 {
	return (x + y) / 2
}

func Product(x, y uint8) uint8 {
	return x * y
}

func Xor(x, y uint8) uint8 {
	return x ^ y
}
