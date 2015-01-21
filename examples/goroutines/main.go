package main

import (
	"fmt"
	"time"
)

func printTime(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", n)
		time.Sleep(200 * time.Millisecond)
		fmt.Println("\n")
	}
}

func main() {
	printTime(2)
	printTime(3)
}
