package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	fmt.Println(os)
	fmt.Print("Go runs on ")

	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
