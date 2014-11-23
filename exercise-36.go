package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for i := range a {
		for j := range a[i] {
			a[i][j] = uint8(0) + uint8(i)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
