package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sliceString := strings.Fields(s)
	num := len(sliceString)
	mountMap := make(map[string]int)

	for i := 0; i < num; i++ {
		(mountMap[sliceString[i]])++
	}

	return mountMap
}

func main() {
	wc.Test(WordCount)
}
