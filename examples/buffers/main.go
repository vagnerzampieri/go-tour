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

	fmt.Println(<-c, <-c, <-c, <-c)
}
