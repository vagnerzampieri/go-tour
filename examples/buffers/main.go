package main

import "fmt"

func add(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func main() {
	c := make(chan int, 3)
	go add(c)
	fmt.Println("channel ->", c)

	for {
		value, ok := <-c
		if ok {
			fmt.Println(value)
		} else {
			break
		}
	}
}
