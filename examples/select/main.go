package main

import "fmt"

func separate(nums []int, o, p chan<- int, ready chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			o <- n
		}
	}
	ready <- true
}

func main() {
	o, p := make(chan int), make(chan int)
	ready := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}
	go separate(nums, o, p, ready)

	var odd, pair []int
	finish := false

	for !finish {
		select {
		case n := <-o:
			odd = append(odd, n)
		case n := <-p:
			pair = append(pair, n)
		case finish = <-ready:
		}
	}

	fmt.Printf("Impares: %v | Pares: %v\n", odd, pair)
}
