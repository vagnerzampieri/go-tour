package main

import "fmt"

type File struct {
	name      string
	length    float64
	caracters int
	words     int
	lines     int
}

func main() {
	file := File{"artigo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(file)
}
