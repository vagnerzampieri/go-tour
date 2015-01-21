package main

import (
	"fmt"
	"time"
)

func sleepGo() {
	fmt.Println("Goroutine sleep 5 seconds ...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine finish.")
}

func main() {
	go sleepGo()

	fmt.Println("Main sleep 3 seconds ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main finish.")
}
