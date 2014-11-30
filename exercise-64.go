package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	fmt.Println("a ->", a)

	c := make(chan int)
	fmt.Println("c ->", c)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("x ->", x)
	fmt.Println("y ->", y)
	fmt.Println("x+y ->", x+y)
}
