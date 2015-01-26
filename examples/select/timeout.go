package main

import (
	"fmt"
	"time"
)

func execute(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}

func main() {
	c := make(chan bool, 1)
	go execute(c)

	fmt.Println("Wait ...")

	finish := false
	for !finish {
		select {
		case finish = <-c:
			fmt.Println("End!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			finish = true
		}
	}
}
