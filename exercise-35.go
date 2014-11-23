package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println("make([]int, 10) ->", pow)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	fmt.Println("make([]int, 10) ->", pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
