package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	expr := regexp.MustCompile("\\b\\w")

	upCase := func(s string) string {
		return strings.ToUpper(s)
	}

	text := "antonio carlos jobim"

	fmt.Println(upCase(text))
	fmt.Println(expr.ReplaceAllStringFunc(text, upCase))
}
