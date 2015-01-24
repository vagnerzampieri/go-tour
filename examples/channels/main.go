package main

import "fmt"

func add(c chan int) {
	c <- 33
}

func main() {
	c := make(chan int)
	fmt.Println("channel ->", c)

	go add(c)

	value := <-c
	fmt.Println(value)
}
