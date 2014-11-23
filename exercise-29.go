package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	// var t *T = new(T)
	v := new(Vertex)
	fmt.Println("init ->", v)

	v.X, v.Y = 11, 9
	fmt.Println("with data ->", v)
}
