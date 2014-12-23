package main

import "fmt"

func main() {
	empty1 := map[int]string{}
	empty2 := make(map[int]string)

	fmt.Println("empty1 ->", empty1)
	fmt.Println("empty2 ->", empty2)
}
